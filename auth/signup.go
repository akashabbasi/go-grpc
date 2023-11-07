package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/akashabbasi/go-grpc/model"
	"github.com/akashabbasi/go-grpc/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrUserAlreadyRegistered = errors.New("User already registered")

func RegisterUser(
	username string, password string,
	name string, collection *mongo.Collection,
	ctx context.Context,
) (*model.UserModel, error) {
	filter := bson.M{"username": username}
	var existingUser model.UserModel

	err := collection.FindOne(ctx, filter).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		pswdHash, err := utils.HashPassword(password)
		if err != nil {
			return nil, fmt.Errorf("unable to hash password: %v", err)
		}
		newUser := model.UserModel{
			Username: username,
			Password: pswdHash,
			Name:     name,
		}

		result, err := collection.InsertOne(ctx, newUser)
		if err != nil {
			return nil, fmt.Errorf("unable to create user: %v", err)
		}

		newUser.ID = result.InsertedID.(primitive.ObjectID)
		return &newUser, nil
	} else if err != nil {
		return nil, fmt.Errorf("unable to create user:" + err.Error())
	}

	return nil, ErrUserAlreadyRegistered
}
