package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-res-gin/models"
	"os"
	"strings"
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
		"name":  user.Name,
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

func (h jwtMiddleware) AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized | token not provided",
		})
		return
	}

	// get token after bearer
	tokenString = strings.Split(tokenString, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNATURE_KEY")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized | token not valid",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println(claims)
		// set claims user id
		c.Set("authId", claims["id"])
		c.Set("authCode", claims["code"])
		c.Set("authName", claims["name"])
		c.Set("authEmail", claims["email"])
		c.Set("authPhone", claims["phone"])
		c.Set("auth", claims)

		// get from c handler

		c.Next()
	} else {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}
}
