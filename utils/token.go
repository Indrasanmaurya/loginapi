package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key") // 🔐 Secret Key (secure रखना future में)

func GenerateJWT(email string) (string, error) {
	// 🧱 Step 1: Token claims banao (data jo token mein jayega)
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(2 * time.Hour).Unix(), // expiry 2 ghante baad
	}

	// 🔐 Step 2: Token create karo with HS256 algo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 🪙 Step 3: Token ko sign karo secret key se
	return token.SignedString(jwtKey)
}
