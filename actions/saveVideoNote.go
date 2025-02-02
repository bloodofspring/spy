package actions

import (
	"fmt"
	"io"
	"log"
	"main/database"
	"net/http"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SaveVideoNoteCallback struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SaveVideoNoteCallback) fabricateAnswer(update tgbotapi.Update, fileID string, videoNoteDuration int) tgbotapi.Chattable {
	filePath := fmt.Sprintf("downloaded_videos/%s.mp4", fileID)
	videoBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Ошибка при чтении файла: %v", err)
		return tgbotapi.NewMessage(update.BusinnesMessage.From.ID, "Не удалось отправить файл")
	}

	videoNoteFile := tgbotapi.FileBytes{
		Name:  "videoNote.mp4",
		Bytes: videoBytes,
	}

	videoNoteMsg := tgbotapi.NewVideoNote(update.BusinnesMessage.From.ID, videoNoteDuration, videoNoteFile)

	defer func() {
		if err := os.Remove(filePath); err != nil {
			log.Printf("Ошибка при удалении файла %s: %v", filePath, err)
		}
	}()

	return videoNoteMsg
}

func (e SaveVideoNoteCallback) Run(update tgbotapi.Update) error {
	if err := database.UpdateAllUserData(update.BusinnesMessage.From.ID, update.BusinnesMessage.BusinessConnectionId, false); err != nil {
		return err
	}

	// Формируем URL для загрузки файла
	fileID := update.BusinnesMessage.ReplyToMessage.VideoNote.FileID
	videoNoteDuration := update.BusinnesMessage.ReplyToMessage.VideoNote.Duration
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

	_, err = e.Client.Send(e.fabricateAnswer(update, fileID, videoNoteDuration))

	return err
}

func (e SaveVideoNoteCallback) GetName() string {
	return e.Name
}
