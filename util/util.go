package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"os"

	"github.com/Arch-4ng3l/GoServerHololens/types"
	jwt "github.com/golang-jwt/jwt/v5"
)

func CalcDistance(obj1, obj2 *types.Object) float32 {
	d := math.Sqrt(math.Pow(float64(obj2.X-obj1.X), 2) + math.Pow(float64(obj2.Z-obj1.Z), 2))
	return float32(d)
}

func CreateJWT(username, password string) (string, error) {
	claims := &jwt.MapClaims{
		"username": username,
		"password": password,
	}
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

}

func CreateHash(s string) string {
	return hex.EncodeToString(sha256.New().Sum([]byte(s)))
}
