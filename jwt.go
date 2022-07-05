package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func (a *Auth) createJWT(expiryDuration time.Duration) string {
	t := time.Now()
	claims := &jwt.RegisteredClaims{
		IssuedAt:  &jwt.NumericDate{Time: t},
		ExpiresAt: &jwt.NumericDate{Time: t.Add(expiryDuration)},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(a.Secret)
	return tokenString
}

func (a *Auth) parseJWT() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v\n", token.Header["alg"])
		}
		return a.Secret, nil
	}
}

func (a *Auth) validateJWT(tokenString string) error {
	_, err := jwt.Parse(tokenString, a.parseJWT())
	if err != nil {
		return err
	} else {
		return nil
	}
}
