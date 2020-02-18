//go:generate protoc -I ../proto ../proto/user.proto --go_out=plugins=grpc,paths=source_relative:../proto

package svc

import (
	"context"

	"github.com/pojntfx/go-support/pkg/controllers"
	"github.com/pojntfx/go-support/pkg/converters"
	"github.com/pojntfx/go-support/pkg/proto"
)

type UserService struct {
	proto.UnimplementedUserServiceServer
	UserController *controllers.UserController
	UserConverter  *converters.UserConverter
}

func (s *UserService) FindById(ctx context.Context, req *proto.UserID) (*proto.User, error) {
	err, internalUser := s.UserController.FindById(int(req.GetID()))
	if err != nil {
		return nil, err
	}

	err, externalUser := s.UserConverter.ToExternal(internalUser)
	return &externalUser, err
}

func (s *UserService) FindAll(ctx context.Context, req *proto.EmptyArgs) (*proto.Users, error) {
	err, internalUsers := s.UserController.FindAll()
	if err != nil {
		return nil, err
	}

	externalUsers := proto.Users{}
	for _, internalUser := range internalUsers {
		err, externalUser := s.UserConverter.ToExternal(internalUser)
		if err != nil {
			return nil, err
		}

		externalUsers.Users = append(externalUsers.GetUsers(), &externalUser)
	}

	return &externalUsers, nil
}
