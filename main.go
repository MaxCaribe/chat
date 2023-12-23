package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/chats", getChats)

	router.Run("localhost:8080")
}

type chat struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func getChats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, chats)
}

var chats = []chat{
	{Id: 1, Title: "Saved messages"},
}
