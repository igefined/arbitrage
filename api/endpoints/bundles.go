package endpoints

import (
	"net/http"

	"github.com/igilgyrg/arbitrage/api/respond"
)

func (e *endpoint) Bundles() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		list, err := e.use.Bundles().List(req.Context())
		if err != nil {
			respond.Error(w, http.StatusInternalServerError, err)
		}
		
		respond.Successfully(w, list)
	}
}
