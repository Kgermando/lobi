package helpers

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/models"
	"github.com/kgermando/lobi/utils" 
)

func ValidToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func ValidUser(id string, p string) bool {
	db := database.DB
	var user models.User
	db.First(&user, id)
	if user.Email == "" {
		return false
	}
	if !utils.CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}
