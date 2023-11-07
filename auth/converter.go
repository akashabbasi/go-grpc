package auth

import (
	"github.com/akashabbasi/go-grpc/model"
	"github.com/akashabbasi/go-grpc/pb"
)

func ConvertUserObjectToUser(model *model.UserModel) *pb.User {
	return &pb.User{
		Username: model.Username,
		Name:     model.Name,
		Id:       int32(model.ID.Timestamp().Day()),
	}
}
