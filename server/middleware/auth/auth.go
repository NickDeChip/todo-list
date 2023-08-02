package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NickDeChip/todo-list/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("jwt")
		userId, err := validateAndReadIdFromJWT(token)
		if err != nil {
			return utility.LogErrorAndSendStatus(c, err, http.StatusUnauthorized)
		}
		c.Locals("UserID", userId)
		return c.Next()
	}
}

func GetUserID(c *fiber.Ctx) (int64, error) {
	userID := c.Locals("UserID")
	if id, ok := userID.(int64); ok {
		return id, nil
	}
	return 0, fmt.Errorf("issue loading userID")
}

func validateAndReadIdFromJWT(token string) (int64, error) {
	tok, err := jwt.Parse(token, func(tok *jwt.Token) (any, error) {
		if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("bad signing method: %v", tok.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := tok.Claims.(jwt.MapClaims); ok && tok.Valid {
		sub := claims["sub"]
		if id, ok := sub.(float64); ok {
			return int64(id), nil
		}
	}
	return 0, fmt.Errorf("could not get ID from token")
}
