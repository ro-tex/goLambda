package lib

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// HasValidAuth extracts and validates the JWT token coming with the request
func HasValidAuth(req *Request) bool {
	if req.Headers["Authorization"] == "" {
		return false
	}

	t := strings.Split(req.Headers["Authorization"], " ")[1]

	return validateJWT(t)
}

func getKey(_ *jwt.Token) (interface{}, error) {
	s, _ := os.LookupEnv("JWT_SECRET")
	return []byte(s), nil
}

func validateJWT(token string) bool {
	t, err := jwt.Parse(token, getKey)
	if err != nil {
		println("Error while parsing JWT token:", err.Error())
		return false
	}
	return t.Valid
}
