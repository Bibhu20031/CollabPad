package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtSecret,
		ContextKey: "user", // Store token info in `user`
	})
}
