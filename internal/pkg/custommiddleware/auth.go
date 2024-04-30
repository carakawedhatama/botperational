package custommiddleware

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func NewAuthMiddleware(secret string, errorHandler interface{}) fiber.Handler {
	handler := jwtware.New(jwtware.Config{
		KeyFunc:      customKeyFunc(secret),
		ErrorHandler: errorHandler.(fiber.ErrorHandler),
	})

	return handler
}

func customKeyFunc(secret string) jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if t.Method.Alg() != jwtware.HS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// Return the key for validation
		return []byte(secret), nil
	}
}
