package models

import pb "github.com/sysradium/petproject/users-api/proto/users/v1"

type User struct {
	Email string
	Login string
	Id    string
}

func (u *User) ToProto() *pb.User {
	return &pb.User{
		Username: u.Login,
		Email:    u.Email,
		Id:       u.Id,
	}
}
