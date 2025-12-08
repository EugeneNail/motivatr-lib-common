package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJsonResponse(handlerFunc HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		status, data := handlerFunc(request)
		if err, isError := data.(error); isError {
			fmt.Printf(err.Error())
			http.Error(writer, err.Error(), status)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(status)

		if status == http.StatusNoContent {
			return
		}

		var buffer bytes.Buffer
		if err := json.NewEncoder(&buffer).Encode(data); err != nil {
			err = fmt.Errorf("encoding response to json: %w", err)
			fmt.Printf(err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err := buffer.WriteTo(writer)
		if err != nil {
			err = fmt.Errorf("writing data from buffer into response writer: %w", err)
			fmt.Printf(err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}
