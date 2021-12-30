package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const AccessToken = 20

type JWT struct {
	secretKey string
}

func CreateToken(username string, duration time.Duration) (string, error) {

	myjwt := JWT{}
	payload, err := NewPayload(username, duration)
	if err != nil {
		fmt.Println("error in payload token:", err)
		fmt.Println("payload token", payload)
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(myjwt.secretKey))

}
