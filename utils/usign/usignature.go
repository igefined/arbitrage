package usign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func SignRequest(queryParam, secretKey string) (hash string) {
	h := hmac.New(sha256.New, []byte(secretKey))

	h.Write([]byte(queryParam))

	hash = hex.EncodeToString(h.Sum(nil))

	return
}
