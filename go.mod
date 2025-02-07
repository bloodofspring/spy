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

require github.com/go-pg/pg/v10 v10.14.0

require (
	github.com/go-pg/zerochecker v0.2.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/bufpool v0.1.11 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.4 // indirect
	github.com/vmihailenco/tagparser v0.1.2 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	mellium.im/sasl v0.3.1 // indirect
)
