package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvaidToken  = errors.New("token is invalid")
)

type Payload struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.RegisteredClaims
}

func NewPayload(id int64, username string, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "gogrpc.com",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Subject:   username,
			Audience:  jwt.ClaimStrings{fmt.Sprint(id)},
		},
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
