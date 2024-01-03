package chats

import (
	"chat/src/common"
	"net/http"
)

func GetChats(writer http.ResponseWriter, request *http.Request) {
	common.JsonResponse(chats, writer)
}

var chats = []Chat{
	{Id: 1, Title: "Saved messages"},
}
