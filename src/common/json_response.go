package common

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(data any, writer http.ResponseWriter) {
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
