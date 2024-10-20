package helpers

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Minute * 60),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return err.Error()
	}
	return signedToken
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errors.New("Sign In to Proceed")
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Sign In to Proceed")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("Sign In to Proceed")
	}

	expClaim, exists := claims["exp"]
	if !exists {
		return nil, errors.New("Expire Claims Is Missing")
	}

	expStr, ok := expClaim.(string)
	if !ok {
		return nil, errors.New("Expire Claims Is not a valid type")
	}

	expTime, err := time.Parse(time.RFC3339, expStr)
	if err != nil {
		return nil, errors.New("Error Parsing Expire Time")
	}

	if time.Now().After(expTime) {
		return nil, errors.New("Token is Expired")
	}
	return token.Claims.(jwt.MapClaims), nil
}
