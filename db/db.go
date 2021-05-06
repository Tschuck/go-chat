package db

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	From    string    `json:"from"`
	Text    string    `json:"text"`
	Created time.Time `json:"created"`
}

var DB *gorm.DB

func GetDb() *gorm.DB {
	if DB != nil {
		return DB
	}

	fmt.Println("Initializing db...")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Message{})

	DB = db

	return DB
}

func LogMessage(message Message) {
	fmt.Printf(
		"%s - %s: %s\n",
		time.Now().Format("January 2, 2006 15:04:05"),
		message.From,
		message.Text,
	)
}

func GetMessages() []Message {
	db := GetDb()
	var messages []Message
	db.Find(&messages)
	return messages
}

func WriteMessage(message Message) {
	GetDb().Create(&message)
}
