package middlewares

import (
	"cleanarchitecture-example/configs"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func GetTokenJWT(headerAuth string) (*Claim, interface{}, string) {
	if headerAuth == "" || !strings.HasPrefix(headerAuth, "Bearer ") {
		return nil, "Invalid or missing authorization header", headerAuth
	}

	jwtSecret := []byte(configs.GetEnv("jwt.secret"))
	headerToken := strings.TrimPrefix(headerAuth, "Bearer ")
	parts := strings.Split(headerToken, ".")
	if len(parts) != 3 {
		return nil, "Invalid token", headerAuth
		// return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		// 	"error": "Invalid Token",
		// })
	}

	claimUser := &Claim{}

	token, err := jwt.ParseWithClaims(headerToken, claimUser, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, "Invalid token", headerToken
	}
	// fmt.Println(token)

	return claimUser, nil, headerToken

}
