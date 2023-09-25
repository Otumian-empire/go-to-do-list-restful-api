package web

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/config"
	"github.com/otumian-empire/go-to-do-list-restful-api/repository"
)

type JWTService interface {
	TokenGenerate(id config.IdType) (string, error)
	TokenValidate(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Id config.IdType `json:"id"`
	jwt.StandardClaims
}

type jwtServices struct {
	privateKey string
	issuer     string
}

func getSecretKey() string {
	return "someHge124237_"
}

func getIssuer() string {
	return "someIssuer12324"
}

func (service *jwtServices) TokenGenerate(id config.IdType) (string, error) {
	claims := &authCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded the web token
	return token.SignedString([]byte(service.privateKey))
}

func (service *jwtServices) TokenValidate(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
		}

		return []byte(service.privateKey), nil
	})
}

func JWTAuthService() JWTService {
	return &jwtServices{
		privateKey: getSecretKey(),
		issuer:     getIssuer(),
	}
}

func AuthorizeJWT(repo repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("AuthorizeJWT")
		apiKey := context.GetHeader("apiKey")

		if err := ValidateApiKey(apiKey); err != nil {
			log.Println(err)
			context.Abort()
			context.JSON(AuthenticationErrorResponse(err.Error()))
			return
		}

		token, err := JWTAuthService().TokenValidate(apiKey)

		if err != nil {
			log.Println(err)
			context.Abort()
			context.JSON(AuthenticationErrorResponse(INVALID_AUTHENTICATION))
			return
		}

		if !token.Valid {
			log.Println("Invalid token")
			context.Abort()
			context.JSON(AuthenticationErrorResponse(INVALID_AUTHENTICATION))
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		log.Println(claims)

		// Check if the token has expired
		// Get the expiration time from the "exp" claim
		expirationTime, ok := claims["exp"].(float64)
		if !ok {
			context.Abort()
			context.JSON(AuthenticationErrorResponse(INVALID_TOKEN_EXPIRATION))
			return
		}

		// Convert the expiration time to a Unix timestamp in seconds
		expirationUnix := int64(expirationTime)

		// Get the current time
		currentTime := time.Now().Unix()

		if currentTime > expirationUnix {
			log.Println(EXPIRED_TOKEN)
			context.Abort()
			context.JSON(AuthenticationErrorResponse(EXPIRED_TOKEN))
			return
		}

		// check the issuer
		issuer, ok := claims["iss"].(string)
		if !ok {
			log.Println(INVALID_TOKEN_ISSUER)
			context.Abort()
			context.JSON(AuthenticationErrorResponse(INVALID_TOKEN_ISSUER))
			return
		}

		if issuer != getIssuer() {
			log.Println(INVALID_TOKEN_ISSUER)
			context.Abort()
			context.JSON(AuthenticationErrorResponse(INVALID_TOKEN_ISSUER))
			return
		}

		// Get the user of this token
		idClaim, ok := claims["id"].(float64)

		log.Println("idClaim", idClaim)
		log.Println("ok", ok)

		if !ok {
			context.Abort()
			context.JSON(AuthenticationErrorResponse(INVALID_AUTHENTICATION))
			return
		}

		userId := config.IdType(idClaim)
		log.Println("userId", userId)

		user, err := repo.ReadUserById(userId)

		if err != nil {
			log.Println(err)
			context.Abort()
			context.JSON(AuthenticationErrorResponse(INVALID_AUTHENTICATION))
			return
		}

		// Remove the password
		user.Password = ""

		context.Set("user", user)

		context.Next()
	}
}
