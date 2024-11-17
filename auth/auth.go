package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("my-secret-key")
var signedToken string

// generates and returns a new signed token, signed using our secretKey from above.
func generateToken(user string) (interface{}, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	var err error
	signedToken, err = token.SignedString(secretKey)

	if err != nil {
		return nil, err
	}

	return signedToken, nil
}

func verifyToken(token string) (bool, error) {
	token = signedToken

	key := secretKey
	verifiedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return false, err
	}

	if !verifiedToken.Valid {
		log.Fatal("Invalid token")
		return false, nil
	}

	return true, nil

}
