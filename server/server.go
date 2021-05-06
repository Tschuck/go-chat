package server

import (
	"encoding/json"
	"fmt"
	"go-decentralized-chat/db"
	"log"
	"net/http"
)

const PORT = "4444"

func onMessage(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(405)
	} else {
		var message db.Message
		err := json.NewDecoder(request.Body).Decode(&message)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		db.LogMessage(message)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(message)

		db.WriteMessage(message)
	}
}

func Start() {
	fmt.Println("Starting messaging server on port " + PORT + ". Listening...")
	http.HandleFunc("/message", onMessage)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
