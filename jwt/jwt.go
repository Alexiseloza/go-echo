package jwt

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateJWT(email string, name string, id string) (string, error) {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic("Error loading .env file")
	}
	myKey := []byte(os.Getenv("SECRET_JWT"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":          email,
		"name":           name,
		"generated_from": "https://www.google.com",
		"id":             id,
		"iat":            time.Now().Unix(),
		"exp":            time.Now().Add(time.Hour * 24).Unix(), //24 Hours
	})

	tokenString, err := token.SignedString(myKey)

	return tokenString, err
}
