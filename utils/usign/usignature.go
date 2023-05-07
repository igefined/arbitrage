package usign

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func SignRequest(queryParam, secretKey string) (hash string) {
	h := hmac.New(sha256.New, []byte(secretKey))

	h.Write([]byte(queryParam))

	hash = hex.EncodeToString(h.Sum(nil))

	return
}

func SignRequestBase64(query, secretKey string) (hash string) {
	var buf bytes.Buffer
	buf.WriteString(query)

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write(buf.Bytes())

	payload := make([]byte, base64.StdEncoding.EncodedLen(buf.Len()))
	base64.StdEncoding.Encode(payload, buf.Bytes())

	hash = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return
}
