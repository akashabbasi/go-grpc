package auth

import (
	"context"
	"errors"

	"github.com/akashabbasi/go-grpc/model"
	"github.com/akashabbasi/go-grpc/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidCreds = errors.New("credentials invalid")
)

func LoginUser(
	username string,
	password string,
	collection *mongo.Collection,
	ctx context.Context,
) (*model.UserModel, error) {
	filter := bson.M{"username": username}
	var user model.UserModel

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	err = utils.CheckPassword(password, user.Password)

	if err != nil {
		return nil, ErrInvalidCreds
	}
	return &user, nil
}
