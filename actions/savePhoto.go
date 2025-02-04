package actions

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SavePhoto struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SavePhoto) fabricateAnswer(update tgbotapi.Update, fileID string) tgbotapi.Chattable {
	filePath := fmt.Sprintf("downloads/%s.jpg", fileID)
	photoBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Ошибка при чтении файла: %v", err)
		return tgbotapi.NewMessage(update.BusinnesMessage.From.ID, "Не удалось отправить файл")
	}

	photoFile := tgbotapi.FileBytes{
		Name:  "photo.jpg",
		Bytes: photoBytes,
	}

	photoMsg := tgbotapi.NewPhoto(update.BusinnesMessage.From.ID, photoFile)

	defer func() {
		if err := os.Remove(filePath); err != nil {
			log.Printf("Ошибка при удалении файла %s: %v", filePath, err)
		}
	}()

	return photoMsg
}

func (e SavePhoto) Run(update tgbotapi.Update) error {
	fileID := update.BusinnesMessage.ReplyToMessage.Photo[len(update.BusinnesMessage.ReplyToMessage.Photo)-1].FileID

	file, err := e.Client.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		return fmt.Errorf("не удалось получить информацию о файле: %v", err)
	}

	photoDir := "downloads"
	if err := os.MkdirAll(photoDir, 0755); err != nil {
		return fmt.Errorf("ошибка создания директории: %w", err)
	}

	fileURL := file.Link(e.Client.Token)

	resp, err := http.Get(fileURL)
	if err != nil {
		return fmt.Errorf("ошибка при скачивании файла: %v", err)
	}
	defer resp.Body.Close()

	outputFile, err := os.Create(fmt.Sprintf("downloads/%s.jpg", fileID))
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	defer outputFile.Close()

	buffer := make([]byte, 32*1024) // 32KB chunks
	var buf bytes.Buffer

	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			// Записываем часть в буфер
			buf.Write(buffer[:n])
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("ошибка при чтении файла: %v", err)
		}
	}

	_, err = buf.WriteTo(outputFile)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %v", err)
	}

	sentMsg, err := e.Client.Send(e.fabricateAnswer(update, fileID))
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(update.BusinnesMessage.From.ID, fmt.Sprintf("Самоуничтожающееся фото от @%s", update.BusinnesMessage.ReplyToMessage.From.UserName))
	msg.ReplyToMessageID = sentMsg.MessageID
	_, err = e.Client.Send(msg)

	return err
}

func (e SavePhoto) GetName() string {
	return e.Name
}