package src

import (
	"chat/src/chats"
	"net/http"
)

func HandleRouting() {
	http.HandleFunc("/chats", chats.GetChats)
}
