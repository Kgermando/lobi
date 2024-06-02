package middleware

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/config"
	"github.com/kgermando/lobi/utils"
)

func IsAuthenticated(c *fiber.Ctx) error {

	cookie := c.Cookies("token")

	if _, err := utils.VerifyJWT(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()
}

func IsTokenExpire(c *fiber.Ctx) error {

	cookie := c.Cookies("token")

	fmt.Println("cookie", cookie)

	if _, err := utils.VerifyJWT(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		fmt.Println("Token expired")
		return c.Next()
	}
	return c.Next()
}

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.ConfigEnv("SECRET"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
