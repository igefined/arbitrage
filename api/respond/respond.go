package respond

import (
	"encoding/json"
	"net/http"
)

func Successfully(writer http.ResponseWriter, data interface{}) {
	Respond(writer, http.StatusOK, data)
}

func Error(writer http.ResponseWriter, code int, err error) {
	Respond(writer, code, err.Error())
}

func Respond(writer http.ResponseWriter, code int, data interface{}) {
	writer.WriteHeader(code)

	if data != nil && data != "" {
		if err := json.NewEncoder(writer).Encode(data); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(err.Error()))

			return
		}
	}
}
