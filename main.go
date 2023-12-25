package main

import (
	"chat/database"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	database.Connection()
	http.HandleFunc("/chats", getChats)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type chat struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func getChats(writer http.ResponseWriter, request *http.Request) {
	toJSON(chats, writer)
}

func toJSON(data any, writer http.ResponseWriter) {
	jsonData, err := json.Marshal(data)
	writer.Header().Set("Content-Type", "application/json")

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errMessage, _ := json.Marshal("Error has occurred")
		writer.Write(errMessage)
		return
	}

	writer.Write(jsonData)
}

var chats = []chat{
	{Id: 1, Title: "Saved messages"},
}
