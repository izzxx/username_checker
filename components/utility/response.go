package utility

import (
	"encoding/json"
	"net/http"

	"github.com/izzxx/username_checker/components/schema"
)

func ResponseApi(w http.ResponseWriter, statusCode int, message, url string) error {
	response := schema.Response{
		StatusCode: statusCode,
		Message:    message,
		Url:        url,
	}

	byteResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, err = w.Write(byteResponse)
	if err != nil {
		return err
	}

	return nil
}
