package web

import "log"

func Recover() {
	if err := recover(); err != nil {
		log.Println(SERVER_RECOVER_FROM_ERROR)
		log.Println(err)
	}
}
