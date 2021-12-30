package common

import (
	Token "main/token"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JwtMiddleWare() echo.MiddlewareFunc {
	var jwt Token.Myjwt
	key := jwt.SecretKey
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &Token.Payload{},
		SigningKey: []byte(key),
	})
}
