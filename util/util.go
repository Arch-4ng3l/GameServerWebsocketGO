package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/Arch-4ng3l/GoServerHololens/types"
	"github.com/TwiN/go-color"
	jwt "github.com/golang-jwt/jwt/v5"
)

const TimeLayout = "15:04:05_2.01.2006"

func PrintSuccess(str string) {
	fmt.Println("[" + time.Now().Format(TimeLayout) + "] " + color.Green + color.Bold + "[*] " + str + color.Reset)
}

func PrintLog(str string) {
	fmt.Println("[" + time.Now().Format(TimeLayout) + "] " + color.Blue + color.Bold + "[*] " + str + color.Reset)
}

func PrintError(err error) {
	fmt.Println("[" + time.Now().Format(TimeLayout) + "] " + color.Red + color.Bold + "[!] " + err.Error() + color.Reset)
}

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
