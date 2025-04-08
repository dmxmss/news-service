package server

import (
  "github.com/gin-gonic/gin"
	"github.com/CherryRadiator/hakathon2025Spring/entities"
	e "github.com/CherryRadiator/hakathon2025Spring/error"
	"github.com/CherryRadiator/hakathon2025Spring/config"
	"github.com/golang-jwt/jwt/v5"

  "net/http"
	"strings"
	"fmt"
	"log"
)

func ErrorCatchMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Next()

    if len(c.Errors) > 0 {
      err := c.Errors.Last().Err
      var error entities.Error
			var statusCode int

      switch err{
			case e.ErrNotAuthorized:
				error.Error = "not authorized"
				statusCode = http.StatusUnauthorized
			case e.ErrDbNewsNotFound:
				error.Error = "news not found"
				statusCode = http.StatusNotFound
			case e.ErrUserIsNotAuthor:
				error.Error = "forbidden"
				statusCode = http.StatusForbidden
			case e.ErrInvalidRequestData:
				error.Error = "invalid request data"
				statusCode = http.StatusBadRequest
      default:
        error.Error = "internal server error"
				statusCode = http.StatusInternalServerError
      }

			c.JSON(statusCode, error)
		}
  }
}

func (s *GinServer) JWTAuthMiddleware(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token not found"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			log.Printf("%s", conf.App.JwtSecret)
			return []byte(conf.App.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			fmt.Print(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		iss, err := token.Claims.GetIssuer()
		c.Set("user_id", iss)

		c.Next()
	}
}
