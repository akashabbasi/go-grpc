package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JwtMaker struct {
	secret string
}

func NewJWTMaker(secret string) (*JwtMaker, error) {
	if len(secret) < minSecretKeySize {
		return nil, fmt.Errorf("invalid secret key size: %d", len(secret))
	}

	return &JwtMaker{
		secret: secret,
	}, nil
}

func (maker *JwtMaker) CreateToken(id int64, username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(id, username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secret))
}

func (maker *JwtMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvaidToken
		}
		return []byte(maker.secret), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(*Payload); !ok || !jwtToken.Valid {
		return nil, ErrInvaidToken
	} else {
		return claims, nil
	}
}
