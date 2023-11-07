package auth

import (
	"time"

	"github.com/akashabbasi/go-grpc/token"
)

func CreateToken(
	tokenMaker token.Maker,
	username string,
	userId int64,
	duration time.Duration,
) (string, error) {
	accessToken, err := tokenMaker.CreateToken(userId, username, duration)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
