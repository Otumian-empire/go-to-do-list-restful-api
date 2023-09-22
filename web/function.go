package web

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Recover() {
	if err := recover(); err != nil {
		log.Println(SERVER_RECOVER_FROM_ERROR)
		log.Println(err)
	}
}

func ConvertStringIdToInt(id string) (int, error) {
	if len(id) < 1 {
		return 0, fmt.Errorf(INVALID_ID)
	}

	intId, intIdErr := strconv.Atoi(id)

	if intIdErr != nil {
		log.Println(intIdErr)
		return 0, fmt.Errorf(INVALID_ID)
	}

	return intId, nil
}

func ValidateString(value string) (validatedString string, isValidString bool) {
	validatedString = strings.Trim(value, " ")
	isValidString = len(validatedString) > 1

	return validatedString, isValidString

}

func convertStringQueryToInt(value string, defaultValue int) int {
	intValue, valueError := strconv.Atoi(value)

	if valueError != nil {
		intValue = defaultValue
	}

	return intValue
}
