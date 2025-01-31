module main

go 1.23.3

// IMPORTANT! add lines to telegram-bot-api/types.go:
// BusinnesMessage *Message `json:"business_message,omitempty"`  (line 50)
// EditedBusinnesMessage *Message `json:"edited_business_message,omitempty"`  (line 51)
// DeletedBusinnesMessage *Message `json:"deleted_business_messages,omitempty"`  (line 52)
require (
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
)
