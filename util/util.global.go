package util

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateId() string {
	new_id := strings.ReplaceAll(uuid.New().String(), "-", "")

	return new_id
}

func HashPassword(text string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(input, hashed string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GenreateJWT(userID, password string) (string, error) {
	claims := jwt.MapClaims{
		"userID":   userID,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
