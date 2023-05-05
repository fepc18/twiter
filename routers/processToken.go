package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fepc18/twiter/bd"
	"github.com/fepc18/twiter/models"
)

var Email string
var IDUser string

// ProcessToken processes the token to extract its values
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("MyKey")
	claims := &models.Claim{}
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}
	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, found, _ := bd.CheckUserExist(claims.Email)
		if found == false {
			return claims, false, string(""), errors.New("user not found")
		}
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}
