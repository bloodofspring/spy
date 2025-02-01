package actions

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SaveVoiceMessageCallback struct {
	Name string
	Client tgbotapi.BotAPI
}


func (e SaveVoiceMessageCallback) fabricateAnswer(update tgbotapi.Update, fileID string) tgbotapi.Chattable {
    filePath := fmt.Sprintf("downloaded_voice/%s.mp3", fileID)
    voiceBytes, err := os.ReadFile(filePath)
    if err != nil {
        log.Printf("Ошибка при чтении файла: %v", err)
        return tgbotapi.NewMessage(update.BusinnesMessage.From.ID, "Не удалось отправить файл")
    }

    voiceNoteFile := tgbotapi.FileBytes{
        Name:  "voice.mp3",
        Bytes: voiceBytes,
    }

    voiceMsg := tgbotapi.NewVoice(update.BusinnesMessage.From.ID, voiceNoteFile)

	defer func() {
		if err := os.Remove(filePath); err != nil {
			log.Printf("Ошибка при удалении файла %s: %v", filePath, err)
		}
	}()

	return voiceMsg
}


func (e SaveVoiceMessageCallback) Run(update tgbotapi.Update) error {
	fileID := update.BusinnesMessage.ReplyToMessage.Voice.FileID
	file, err := e.Client.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		return fmt.Errorf("ошибка получения информации о файле: %w", err)
	}

	fileURL := file.Link(e.Client.Token)

	// Создаем директорию для сохранения, если её нет
	voiceDir := "downloaded_voice"
	if err := os.MkdirAll(voiceDir, 0755); err != nil {
		return fmt.Errorf("ошибка создания директории: %w", err)
	}

	// Формируем путь для сохранения файла
	filePath := filepath.Join(voiceDir, fmt.Sprintf("%s.mp3", fileID))

	// Скачиваем файл частями
	client := &http.Client{}
	resp, err := client.Get(fileURL)
	if err != nil {
		return fmt.Errorf("ошибка при получении файла: %w", err)
	}
	defer resp.Body.Close()

	// Создаем файл для записи
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer out.Close()

	// Копируем данные частями, используя буфер
	buffer := make([]byte, 32*1024) // 32KB буфер
	_, err = io.CopyBuffer(out, resp.Body, buffer)
	if err != nil {
		return fmt.Errorf("ошибка сохранения файла: %w", err)
	}

	// Отправляем сообщение об успешном сохранении
	_, err = e.Client.Send(e.fabricateAnswer(update, fileID))

	return err
}

func (e SaveVoiceMessageCallback) GetName() string {
	return e.Name
}
