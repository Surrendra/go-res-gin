package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-res-gin/models"
	"os"
	"time"
)

type jwtMiddleware struct {
}

func NewJwtMiddleware() *jwtMiddleware {
	return &jwtMiddleware{}
}

type JwtMiddleware interface {
	GenerateJWTToken(user models.User) (string, error)
}

func (h jwtMiddleware) GenerateJWTToken(user models.User) (string, error) {
	var claims = jwt.MapClaims{
		"id":    user.Id,
		"code":  user.Code,
		"email": user.Email,
		"phone": user.Phone,
		"exp":   time.Now().Add(time.Hour * 34).Unix(),
		"iat":   time.Now().Unix(),
		"type":  "JWT",
	}
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("JWT_SIGNATURE_KEY")
	signToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return signToken, nil
}
