package hmac

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"strings"
	"time"
)

const (
	veracodeRequestVersionString = "vcode_request_version_1"
	dataFormat                   = "id=%s&host=%s&url=%s&method=%s"
	headerFormat                 = "%s id=%s,ts=%s,nonce=%X,sig=%X"
	veracodeHMACSHA256           = "VERACODE-HMAC-SHA-256"
)

func CurrentTimestamp() int64 {
	return time.Now().UnixMilli()
}

func GenerateNonce(size int) ([]byte, error) {
	nonce := make([]byte, size)
	_, err := rand.Read(nonce)

	if err != nil {
		return nil, err
	}

	return nonce, nil
}

func hmac256(message, key []byte) []byte {
	sha := hmac.New(sha256.New, key)
	sha.Write(message)
	return sha.Sum(nil)
}

func removeRegion(apiCredential string) string {
	return strings.Split(apiCredential, "-")[1]
}

func calculateSignature(key, nonce, timestamp, data []byte) []byte {
	encryptedNonce := hmac256(nonce, key)
	encryptedTimestamp := hmac256(timestamp, encryptedNonce)
	signingKey := hmac256([]byte(veracodeRequestVersionString), encryptedTimestamp)
	return hmac256(data, signingKey)
}
