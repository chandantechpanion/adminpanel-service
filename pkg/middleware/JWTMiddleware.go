package middleware

import (
	"errors"
	"fmt"
	"github.com/chandantechpanion/adminpanel-service/pkg/jwtService"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) > 0 {
			tokenString := authHeader[len(BEARER_SCHEMA)+1:]
			token, err := jwtService.NewJWTAuthService().ValidateToken(tokenString)
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				fmt.Println(claims)
			} else {
				fmt.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			c.AbortWithError(http.StatusUnauthorized, errors.New("please pass Authorization header"))
		}
	}
}
