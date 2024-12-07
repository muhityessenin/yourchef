package interfaces

import "user_service/internal/domain/user"

type UserService interface {
	GetUsers() ([]user.Entity, error)
	CreateUser(input *user.InputResponse) (int, error)
	Login(email, password string) (int, error)
	GetUserById(id string) (user.Entity, error)
	UpdateUser(id string, input *user.InputResponse) (int, error)
	DeleteUser(id string) (int, error)
	SearchUser(queryType, query string) ([]user.Entity, error)
}
