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

type SaveVideoMessageCallback struct {
	Name string
	Client tgbotapi.BotAPI
}


func (e SaveVideoMessageCallback) fabricateAnswer(update tgbotapi.Update, fileID string) tgbotapi.Chattable {
    filePath := fmt.Sprintf("downloaded_videos/%s.mp4", fileID)
    photoBytes, err := os.ReadFile(filePath)
    if err != nil {
        log.Printf("Ошибка при чтении файла: %v", err)
        return tgbotapi.NewMessage(update.BusinnesMessage.From.ID, "Не удалось отправить файл")
    }

    videoNoteFile := tgbotapi.FileBytes{
        Name:  "videoNote.jpg",
        Bytes: photoBytes,
    }

    videoNoteMsg := tgbotapi.NewVideoNote(update.BusinnesMessage.From.ID, 60, videoNoteFile)

	defer func() {
		if err := os.Remove(filePath); err != nil {
			log.Printf("Ошибка при удалении файла %s: %v", filePath, err)
		}
	}()

	return videoNoteMsg
}


func (e SaveVideoMessageCallback) Run(update tgbotapi.Update) error {
	// Формируем URL для загрузки файла
	fileID := update.BusinnesMessage.ReplyToMessage.VideoNote.FileID
	file, err := e.Client.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		return fmt.Errorf("не удалось получить информацию о файле: %v", err)
	}

	fileURL := file.Link(e.Client.Token)

	// Создаем HTTP клиент
	client := &http.Client{}
	resp, err := client.Get(fileURL)
	if err != nil {
		return fmt.Errorf("ошибка при получении файла: %w", err)
	}
	defer resp.Body.Close()

	// Создаем директорию для сохранения, если её нет
	saveDir := "downloaded_videos"
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return fmt.Errorf("ошибка при создании директории: %w", err)
	}

	// Создаем файл для сохранения
	filePath := filepath.Join(saveDir, fileID+".mp4")
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %w", err)
	}
	defer out.Close()

	// Используем буферизированное копирование для больших файлов
	buf := make([]byte, 1024*1024) // 1MB буфер
	_, err = io.CopyBuffer(out, resp.Body, buf)
	if err != nil {
		return fmt.Errorf("ошибка при сохранении файла: %w", err)
	}

	_, err = e.Client.Send(e.fabricateAnswer(update, fileID))

	return err
}


func (e SaveVideoMessageCallback) GetName() string {
	return e.Name
}
