package actions

import (
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
	// msg := tgbotapi.NewPhoto(update.BusinnesMessage.From.ID, tgbotapi.FileID(update.BusinnesMessage.ReplyToMessage.Photo[0].FileID))
	// msg.Caption = "Фото собеседника"
	msg := tgbotapi.NewMessage(update.BusinnesMessage.From.ID, "test")

	return msg
}

func (e SaveFile) Run(update tgbotapi.Update) error {
	file, err := e.Client.GetFile(tgbotapi.FileConfig{FileID: update.BusinnesMessage.ReplyToMessage.Photo[0].FileID})
    if err != nil {
        log.Fatal("Something went wrong with saving user file!")
    }

	fmt.Println(file.Link(e.Client.Token))

	testLink := fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", e.Client.Token, update.BusinnesMessage.ReplyToMessage.Photo[0].FileID)
	fmt.Println(testLink)

    resp, err := http.Get(file.Link(e.Client.Token))
    if err != nil {
        log.Fatal("Something went wrong with accessing user file url!")
    }

    defer resp.Body.Close()

    // if _, err := os.Stat("/downloads"); err == nil {
    //  // path/to/whatever exists
    //   } else if errors.Is(err, os.ErrNotExist) {
    //  os.Mkdir("/downloads", os.FileMode{}})
    //   }

    out, err := os.Create("downloads/test.jpg")
    if err != nil {
        log.Fatal(err)
    }
    
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    if err != nil {
        log.Fatal(err)
    }

	if _, err := e.Client.Send(e.FabricateAnswer(update)); err != nil {
		return err
	}

	return nil
}

func (e SaveFile) GetName() string {
	return e.Name
}


