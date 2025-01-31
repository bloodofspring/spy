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


type SaveFile struct {
	Name   string
	Client tgbotapi.BotAPI
}

func (e SaveFile) FabricateAnswer(update tgbotapi.Update) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(update.BusinnesMessage.From.ID, "Сохранено горячее фото собеседника!")

	return msg
}

func (e SaveFile) Run(update tgbotapi.Update) error {
	fileID := update.BusinnesMessage.ReplyToMessage.Photo[len(update.BusinnesMessage.ReplyToMessage.Photo)-1].FileID
	
	file, err := e.Client.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		return fmt.Errorf("не удалось получить информацию о файле: %v", err)
	}

	fileURL := file.Link(e.Client.Token)

	resp, err := http.Get(fileURL)
	if err != nil {
		return fmt.Errorf("ошибка при скачивании файла: %v", err)
	}
	defer resp.Body.Close()

	// Создаем файл для сохранения
	outputFile, err := os.Create(fmt.Sprintf("downloads/%s.jpg", fileID))
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	defer outputFile.Close()

	// Буфер для чтения по частям
	buffer := make([]byte, 32*1024) // 32KB chunks
	
	// Создаем буфер в памяти для накопления данных
	var buf bytes.Buffer
	
	// Читаем файл частями
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

	// Записываем весь буфер в файл
	_, err = buf.WriteTo(outputFile)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %v", err)
	}

	log.Printf("Файл успешно сохранен: %s.jpg", fileID)
	_, err = e.Client.Send(e.FabricateAnswer(update))

	return err
}

func (e SaveFile) GetName() string {
	return e.Name
}


