package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func JwtToken(username string) string {
	expirationTime := time.Now().Add(time.Duration(TokenExpireTime) * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		ThrowException(INTERNAL_SERVER_ERROR)
	}
	return tokenString
}

func JwtDecode(tokenStr string) string {
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ThrowException(UNAUTHORIZED)
		}
		ThrowException(UNAUTHORIZED)
	}
	if !tkn.Valid {
		ThrowException(UNAUTHORIZED)
	}
	return claims.Username
}
