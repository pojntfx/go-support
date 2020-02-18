package converters

import (
	"github.com/pojntfx/go-support/pkg/models"
	"github.com/pojntfx/go-support/pkg/proto"
)

type UserConverter struct {
}

func (u *UserConverter) ToExternal(internal models.UserModel) (error, proto.User) {
	return nil, proto.User{
		ID:         int64(internal.ID),
		Email:      internal.Email,
		FirstName:  internal.FirstName,
		SecondName: internal.SecondName,
		UserName:   internal.UserName,
	}
}
