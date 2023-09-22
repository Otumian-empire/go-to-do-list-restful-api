package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/otumian-empire/go-to-do-list-restful-api/config"
	"github.com/otumian-empire/go-to-do-list-restful-api/repository"
	"github.com/otumian-empire/go-to-do-list-restful-api/web"
)

func main() {
	defer web.Recover()

	ENV_CONST, err := config.GetEnvirons()

	if err != nil {
		log.Printf(err.Error())
		log.Fatalln(web.SERVER_LOADING_CREDENTIALS_ERROR)
	}

	dbConfig := mysql.Config{
		User:                 ENV_CONST.DatabaseUsername,
		Passwd:               ENV_CONST.DatabasePassword,
		Net:                  "tcp", // is tcp by default
		Addr:                 fmt.Sprintf("%s:%s", ENV_CONST.DatabaseHost, ENV_CONST.DatabasePort),
		DBName:               ENV_CONST.DatabaseName,
		AllowNativePasswords: true,
	}

	store, err := repository.NewRepository(ENV_CONST.DatabaseDriverName, dbConfig.FormatDSN())

	if err != nil {
		log.Printf(err.Error())
		log.Fatalln(err)
	}

	log.Println(web.DATABASE_CONNECTED)

	// This handler here is a not a handler as defined in the NewHandler
	// It is the route on passed to the new handler that is returned
	handler := web.NewHandler(*store, gin.Default())

	log.Println(fmt.Sprintf("%v: %v", web.SERVER_RUNNING_ON_PORT, ENV_CONST.ServerPort))
	http.ListenAndServe(fmt.Sprintf(":%v", ENV_CONST.ServerPort), handler)
}
