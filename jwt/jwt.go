package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fepc18/twiter/models"
)

// GenerateJWT generates the encryption of the token
func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("MyKey")
	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastname":  t.LastName,
		"birthdate": t.BirthDate,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
