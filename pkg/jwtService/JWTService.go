package jwtService

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//jwt service
type Service interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	UserId string `json:"user_id"`
	User   bool   `json:"user"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//auth-jwt
func NewJWTAuthService() Service {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Techpanion",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "Techpanion.comForSupplierPortal"
	}
	return secret
}

func (service *jwtServices) GenerateToken(userId string, isUser bool) string {
	claims := &authCustomClaims{
		userId,
		isUser,
		"Admin",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
