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

type SaveVoiceMessage struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SaveVoiceMessage) fabricateAnswer(update tgbotapi.Update, fileID string) tgbotapi.Chattable {
	filePath := fmt.Sprintf("downloads/%s.mp3", fileID)
	voiceBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("ошибка при чтении файла: %v", err)
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

func (e SaveVoiceMessage) Run(update tgbotapi.Update) error {
	if err := database.UpdateAllUserData(update.BusinnesMessage.From.ID, update.BusinnesMessage.BusinessConnectionId, false); err != nil {
		return err
	}

	fileID := update.BusinnesMessage.ReplyToMessage.Voice.FileID
	file, err := e.Client.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		return fmt.Errorf("ошибка получения информации о файле: %w", err)
	}

	fileURL := file.Link(e.Client.Token)

	voiceDir := "downloads"
	if err := os.MkdirAll(voiceDir, 0755); err != nil {
		return fmt.Errorf("ошибка создания директории: %w", err)
	}

	filePath := filepath.Join(voiceDir, fmt.Sprintf("%s.mp3", fileID))

	client := &http.Client{}
	resp, err := client.Get(fileURL)
	if err != nil {
		return fmt.Errorf("ошибка при получении файла: %w", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer out.Close()

	buffer := make([]byte, 32*1024)
	_, err = io.CopyBuffer(out, resp.Body, buffer)
	if err != nil {
		return fmt.Errorf("ошибка сохранения файла: %w", err)
	}

	sentMsg, err := e.Client.Send(e.fabricateAnswer(update, fileID))
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(update.BusinnesMessage.From.ID, fmt.Sprintf("Самоуничтожающееся голосовое сообщение от @%s", update.BusinnesMessage.ReplyToMessage.From.UserName))
	msg.ReplyToMessageID = sentMsg.MessageID
	_, err = e.Client.Send(msg)

	return err
}

func (e SaveVoiceMessage) GetName() string {
	return e.Name
}
