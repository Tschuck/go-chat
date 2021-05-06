package main

import (
	"fmt"
	"go-decentralized-chat/db"
	"go-decentralized-chat/server"
	"time"
)

func logHeader() {
	fmt.Println("--------------------------------------")
	fmt.Println("| DECENTRALIZED - CHAT - GO              |")
	fmt.Printf("|    SERVER: listening on PORT %s  |\n", server.PORT)
	fmt.Println("|    CLIENT: waiting for user inputs |")
	fmt.Println("--------------------------------------\n\n")
}

func main() {
	logHeader()
	db.GetDb()

	fmt.Println("\n\nprevious messages:")

	messages := db.GetMessages()

	for i := 0; i < len(messages); i++ {
		db.LogMessage(messages[i])
	}

	go server.Start()
	fmt.Println("test")
	time.Sleep(100000 * time.Millisecond)
}
