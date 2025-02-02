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

type SaveVideoMessage struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SaveVideoMessage) fabricateAnswer(update tgbotapi.Update, fileID string) tgbotapi.Chattable {
	filePath := fmt.Sprintf("downloads/%s.mp4", fileID)
	videoBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Ошибка при чтении файла: %v", err)
		return tgbotapi.NewMessage(update.BusinnesMessage.From.ID, "Не удалось отправить файл")
	}

	videoFile := tgbotapi.FileBytes{
		Name:  "video.mp4",
		Bytes: videoBytes,
	}

	videoMsg := tgbotapi.NewVideo(update.BusinnesMessage.From.ID, videoFile)

	defer func() {
		if err := os.Remove(filePath); err != nil {
			log.Printf("Ошибка при удалении файла %s: %v", filePath, err)
		}
	}()

	return videoMsg
}

func (e SaveVideoMessage) Run(update tgbotapi.Update) error {
	if err := database.UpdateAllUserData(update.BusinnesMessage.From.ID, update.BusinnesMessage.BusinessConnectionId, false); err != nil {
		return err
	}

	fileID := update.BusinnesMessage.ReplyToMessage.Video.FileID
	file, err := e.Client.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		return fmt.Errorf("не удалось получить информацию о файле: %v", err)
	}

	fileURL := file.Link(e.Client.Token)

	client := &http.Client{}
	resp, err := client.Get(fileURL)
	if err != nil {
		return fmt.Errorf("ошибка при получении файла: %w", err)
	}
	defer resp.Body.Close()

	saveDir := "downloads"
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return fmt.Errorf("ошибка при создании директории: %w", err)
	}

	filePath := filepath.Join(saveDir, fileID+".mp4")
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %w", err)
	}
	defer out.Close()

	buf := make([]byte, 1024*1024) // 1MB буфер
	_, err = io.CopyBuffer(out, resp.Body, buf)
	if err != nil {
		return fmt.Errorf("ошибка при сохранении файла: %w", err)
	}

	sentMsg, err := e.Client.Send(e.fabricateAnswer(update, fileID))
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(update.BusinnesMessage.From.ID, fmt.Sprintf("Самоуничтожающееся видео от @%s", update.BusinnesMessage.ReplyToMessage.From.UserName))
	msg.ReplyToMessageID = sentMsg.MessageID
	_, err = e.Client.Send(msg)

	return err
}

func (e SaveVideoMessage) GetName() string {
	return e.Name
}