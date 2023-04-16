package middleware

import (
	"errors"
	"net/http"

	"github.com/igilgyrg/arbitrage/api/respond"
	"github.com/igilgyrg/arbitrage/log"
)

var (
	ErrMissingAuthorizationKey = errors.New("missing authorization key in request header X-API-Key")
	ErrAccessDenied            = errors.New("access denied")
)

func ClientAccess(log *log.Logger, validKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKeyFromHeader, err := ApiKeyExtractor(r)
			if err != nil {
				log.Warn(err)
				respond.Error(w, http.StatusForbidden, err)

				return
			}

			if apiKeyFromHeader != validKey {
				respond.Error(w, http.StatusForbidden, ErrAccessDenied)

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func ApiKeyExtractor(r *http.Request) (string, error) {
	const authKey = "X-API-Key"

	keyFromHeader := r.Header.Get(authKey)
	if len(keyFromHeader) == 0 {
		return "", ErrMissingAuthorizationKey
	}

	return keyFromHeader, nil
}
