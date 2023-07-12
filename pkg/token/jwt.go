package token

import (
	"fmt"
	"time"

	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"github.com/Najwan160/go-experiment-3/code/domain/model"
	"github.com/dgrijalva/jwt-go/v4"
)

type customClaims struct {
	model.Auth
	jwt.StandardClaims
}

type JWT struct {
	ttl          int
	signatureKey string
}

func NewJWT(ttl int, signatureKey string) base.Token {
	return &JWT{
		ttl:          ttl,
		signatureKey: signatureKey,
	}
}

// LOGIN
func (j *JWT) Generate(a model.Auth) (string, error) {
	expiredAt, err := jwt.ParseTime(time.Now().Add(time.Hour * 24 * time.Duration(j.ttl)).Unix())
	if err != nil {
		return "", err
	}

	claims := customClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "go-experiment",
			ExpiresAt: expiredAt,
		},
		Auth: a,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.signatureKey))
}

func (j *JWT) Parse(tokenString string) (a model.Auth, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.signatureKey), nil
	})

	if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
		return claims.Auth, nil
	}

	return model.Auth{}, err
}
