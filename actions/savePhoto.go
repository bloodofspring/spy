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

func (e SaveFile) fabricateAnswer(update tgbotapi.Update, fileID string) tgbotapi.Chattable {
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

func (e SaveFile) Run(update tgbotapi.Update) error {
	fmt.Printf("\n\n%v\n", update)

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
	_, err = e.Client.Send(e.fabricateAnswer(update, fileID))

	return err
}

func (e SaveFile) GetName() string {
	return e.Name
}


