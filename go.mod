module main

go 1.23.3

// IMPORTANT! add line "BusinnesMessage *Message `json:"business_message,omitempty"`" to line 50 in telegram-bot-api/types.go
require (
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
)
