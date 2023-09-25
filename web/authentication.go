package web

import (
	"fmt"
	"log"
	"strconv"
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

func (service *jwtServices) TokenGenerate(id config.IdType) (string, error) {
	env, _ := config.GetEnvirons()

	ttl, _ := strconv.Atoi(env.JwtTtl)

	log.Println("ttl", ttl)

	claims := &authCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(ttl)).Unix(),
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
	env, _ := config.GetEnvirons()
	return &jwtServices{
		privateKey: env.JwtSecret,
		issuer:     env.JwtIssuer,
	}
}

func IsExpiredToken(claims jwt.MapClaims) bool {
	// Get the expiration time from the "exp" claim
	expirationTime, ok := claims["exp"].(float64)
	if !ok {
		return true
	}

	// Convert the expiration time to a Unix timestamp in seconds
	expirationUnix := int64(expirationTime)

	// Get the current time
	currentTime := time.Now().Unix()

	log.Println(currentTime, currentTime > expirationUnix)

	return currentTime > expirationUnix
}

func IsInvalidIssuer(claims jwt.MapClaims, issuer string) bool {
	// check the issuer
	issuer, ok := claims["iss"].(string)
	if !ok {
		log.Println(INVALID_TOKEN_ISSUER)
		return true
	}

	return issuer != issuer
}

func GetIdFromClaim(claims jwt.MapClaims) (config.IdType, error) {
	idClaim, ok := claims["id"].(float64)

	if !ok {
		return 0, fmt.Errorf(INVALID_TOKEN_SUBJECT_ID)
	}

	return config.IdType(idClaim), nil
}

func AuthorizeJWT(repo repository.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		env, _ := config.GetEnvirons()
		log.Println("env", env)

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
			log.Println("token.Valid ")
			context.Abort()
			context.JSON(AuthenticationErrorResponse(INVALID_AUTHENTICATION))
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		// Check if the token has expired
		if IsExpiredToken(claims) {
			context.Abort()
			context.JSON(AuthenticationErrorResponse(INVALID_TOKEN_EXPIRATION))
			return
		}

		// Check the issuer
		if IsInvalidIssuer(claims, env.JwtIssuer) {
			context.Abort()
			context.JSON(AuthenticationErrorResponse(INVALID_TOKEN_ISSUER))
			return
		}

		// Get the user of this token
		userId, idErr := GetIdFromClaim(claims)
		if idErr != nil {
			log.Println(idErr)
			context.Abort()
			context.JSON(AuthenticationErrorResponse(INVALID_AUTHENTICATION))
			return
		}

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
