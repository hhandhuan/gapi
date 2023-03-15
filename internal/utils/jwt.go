package utils

import (
	"errors"
	"time"
	"zhengze/pkg/conf"

	"github.com/golang-jwt/jwt/v4"
)

type Jwt struct {
	conf *conf.Jwt
}

func NewJwt(conf *conf.Jwt) *Jwt {
	return &Jwt{conf: conf}
}

func (j *Jwt) CreateToken(sub string) (string, error) {
	obj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    j.conf.Issuer,
		Subject:   sub,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.conf.Ttl) * time.Second)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	return obj.SignedString([]byte(j.conf.Secret))
}

func (j *Jwt) ParseJwtToken(token string) (*jwt.RegisteredClaims, error) {
	obj, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(j.conf.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := obj.Claims.(*jwt.RegisteredClaims)
	if ok && obj.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
