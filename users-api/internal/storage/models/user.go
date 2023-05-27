package models

import (
	"github.com/google/uuid"
	pb "github.com/sysradium/petproject/users-api/api/users/v1"
)

type User struct {
	Email string
	Login string
	Id    uuid.UUID
	RowId int
}

func (u *User) ToProto() *pb.User {
	return &pb.User{
		Username: u.Login,
		Email:    u.Email,
		Id:       u.Id.String(),
	}
}
