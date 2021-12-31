package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const AccessToken = 100000000000

type Myjwt struct {
	SecretKey string
}

func CreateToken(username string, duration time.Duration) (string, error) {

	myjwt := Myjwt{}
	payload, err := NewPayload(username, duration)
	if err != nil {
		fmt.Println("error in payload token:", err)
		fmt.Println("payload token", payload)
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(myjwt.SecretKey))
}
