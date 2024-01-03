package common

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"io"
	"math"
	"net/http/httptest"
	"testing"
)

func TestJsonResponse(t *testing.T) {
	var testableJson testJson
	gofakeit.Struct(&testableJson)

	recorder := httptest.NewRecorder()
	JsonResponse(testableJson, recorder)
	responseResult := recorder.Result()
	if responseResult.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Must contain json content-type header")
	}

	read, _ := io.ReadAll(responseResult.Body)
	newJson := new(testJson)
	json.Unmarshal(read, &newJson)

	if newJson.Age != testableJson.Age || newJson.Name != testableJson.Name {
		t.Errorf("Data is serialized with mutations")
	}
}

func TestJsonResponseErrorHandling(t *testing.T) {
	recorder := httptest.NewRecorder()
	JsonResponse(math.Inf(1), recorder)
	responseResult := recorder.Result()
	if responseResult.StatusCode != 500 {
		t.Errorf("Must have status 500")
	}

	read, _ := io.ReadAll(responseResult.Body)
	newJson := ""
	json.Unmarshal(read, &newJson)

	if newJson == "" {
		t.Errorf("Must contain error message")
	}
}

type testJson struct {
	Name string `fake:"{firstname}"`
	Age  int    `fake:"{number:1,80}"`
}
