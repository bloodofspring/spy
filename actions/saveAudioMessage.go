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

type SaveAudioMessageCallback struct {
	Name string
	Client *tgbotapi.BotAPI
}


// ToDo: redo for sending audios
func (e *SaveAudioMessageCallback) fabricateAnswer(update tgbotapi.Update, fileID string) tgbotapi.Chattable {
    filePath := fmt.Sprintf("downloaded_videos/%s.mp3", fileID)
    audioBytes, err := os.ReadFile(filePath)
    if err != nil {
        log.Printf("Ошибка при чтении файла: %v", err)
        return tgbotapi.NewMessage(update.BusinnesMessage.From.ID, "Не удалось отправить файл")
    }

    audioNoteFile := tgbotapi.FileBytes{
        Name:  "audio.mp3",
        Bytes: audioBytes,
    }

    audioMsg := tgbotapi.NewAudio(update.BusinnesMessage.From.ID, audioNoteFile)

	defer func() {
		if err := os.Remove(filePath); err != nil {
			log.Printf("Ошибка при удалении файла %s: %v", filePath, err)
		}
	}()

	return audioMsg
}


func (e *SaveAudioMessageCallback) Handle(update tgbotapi.Update) error {
	fileID := update.Message.ReplyToMessage.Audio.FileID
	file, err := e.Client.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		return fmt.Errorf("ошибка получения информации о файле: %w", err)
	}

	fileURL := file.Link(e.Client.Token)

	// Создаем директорию для сохранения, если её нет
	audioDir := "downloaded_audio"
	if err := os.MkdirAll(audioDir, 0755); err != nil {
		return fmt.Errorf("ошибка создания директории: %w", err)
	}

	// Формируем путь для сохранения файла
	filePath := filepath.Join(audioDir, fmt.Sprintf("%s.mp3", fileID))

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

func (e *SaveAudioMessageCallback) GetName() string {
	return e.Name
}
