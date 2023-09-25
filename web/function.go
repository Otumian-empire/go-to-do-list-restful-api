package web

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

func Recover() {
	if err := recover(); err != nil {
		log.Println(SERVER_RECOVER_FROM_ERROR)
		log.Println(err)
	}
}

func ValidateApiKey(token string) error {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		log.Println(INVALID_TOKEN_FORMAT)
		return fmt.Errorf(INVALID_TOKEN_FORMAT)
	}

	_, headerErr := base64.RawURLEncoding.DecodeString(parts[0])
	if headerErr != nil {
		log.Println(INVALID_TOKEN_HEADER, headerErr)
		return fmt.Errorf(INVALID_TOKEN_HEADER)
	}

	_, payloadErr := base64.RawURLEncoding.DecodeString(parts[1])
	if payloadErr != nil {
		log.Println(INVALID_TOKEN_PAYLOAD, payloadErr)
		return fmt.Errorf(INVALID_TOKEN_PAYLOAD)
	}

	return nil
}
