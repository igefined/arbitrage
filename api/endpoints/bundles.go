package endpoints

import (
	"net/http"

	"github.com/igilgyrg/arbitrage/api/respond"
)

func (e *endpoint) Bundles() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		respond.Successfully(writer, "bundles")
	}
}
