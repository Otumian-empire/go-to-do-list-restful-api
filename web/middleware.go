package web

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/model"
	"github.com/otumian-empire/go-to-do-list-restful-api/repository"
)

func ApiKeyAuth(_repository repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println(context.Request.Header)
		var apiKey = context.Request.Header.Get("apiKey")

		if len(apiKey) < 1 {
			log.Println("There is no api-key")
			log.Println(INVALID_AUTHENTICATION)
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			context.Abort()
			return
		}

		user, err := _repository.User.ReadUserByUsername(apiKey)

		if err != nil {
			log.Println("There is no api-key")
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			context.Abort()
			return
		}

		context.Set("user", model.User{
			Id:        user.Id,
			Username:  user.Username,
			Password:  "",
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})

		context.Next()
	}
}

// TODO: Use jwt in place of apiKey authentication
