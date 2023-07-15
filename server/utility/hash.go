package utility

import (
	"crypto/sha256"
	"os"
)

func HashPassword(password string) []byte {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hash.Sum([]byte(os.Getenv("SALT")))
}
