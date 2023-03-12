package endpoints

import (
	"net/http"

	"github.com/igilgyrg/arbitrage/api/respond"
)

func (e *endpoint) Status() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ok := struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		}

		respond.Successfully(writer, ok)
	}
}
