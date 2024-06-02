package helpers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kgermando/lobi/config"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/models"
)

func UserJWT(c *fiber.Ctx) (models.User, bool) {
	var isAuth bool
	var user models.User

	getToken := c.Cookies("token")
	fmt.Println("getToken:", getToken)

	if getToken != "" {
		token, err := jwt.ParseWithClaims(getToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.ConfigEnv("SECRET")), nil //using the SecretKey which was generated in th Login function
		})
		if err != nil {
			fmt.Println("StatusUnauthorized", fiber.StatusUnauthorized)
			isAuth = false
			user = models.User{}
			return user, isAuth
		}
		if !token.Valid {
			fmt.Println("token not Valid")
			isAuth = false
			user = models.User{}
			return user, isAuth
		}

		claims := token.Claims.(*jwt.RegisteredClaims)

		database.DB.Where("id = ?", claims.Issuer).First(&user)
		if user.Email == "" {
			isAuth = false
		} else {
			isAuth = true
		}
		return user, isAuth
	} else {
		isAuth = false
		user = models.User{}
		return user, isAuth
	}

}

// func UserJWT1(c *fiber.Ctx) (models.User, bool, error) {
// 	var isAuth bool
// 	var user models.User

// 	clientIP := c.IP()
// 	fmt.Println("Client IP:", clientIP)

// 	getToken, err := database.RedisClient.Get(context.Background(), clientIP).Result()
// 	if err != nil {
// 		if err == redis.Nil {
// 			// Ignore nil value
// 			fmt.Println("Key not found")
// 			isAuth = false
// 			user = models.User{}
// 			return user, false, c.Redirect("/web")
// 		}
// 		// Handle other errors
// 		return models.User{}, false, err
// 	}

// 	fmt.Println("getToken:", getToken)

// 	if getToken != "" {
// 		token, err := jwt.ParseWithClaims(getToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
// 			return []byte(config.ConfigEnv("SECRET")), nil //using the SecretKey which was generated in th Login function
// 		})

// 		if err != nil {
// 			fmt.Println("StatusUnauthorized", fiber.StatusUnauthorized)
// 			isAuth = false
// 			user = models.User{}
// 			return user, false, c.Redirect("/web")
// 		}

// 		claims := token.Claims.(*jwt.RegisteredClaims)

// 		database.DB.Where("id = ?", claims.Issuer).First(&user)
// 		if user.Email == "" {
// 			isAuth = false
// 		} else {
// 			isAuth = true
// 		}
// 		return user, isAuth, err

// 	} else {
// 		isAuth = false
// 		user = models.User{}
// 		return user, isAuth, err
// 	}

// }
