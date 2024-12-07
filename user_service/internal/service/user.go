package service

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"user_service/internal/domain/user"
	interf "user_service/internal/repository/interfaces"
	services "user_service/internal/service/interfaces"
)

type UserService struct {
	UserRepository interf.UserRepository
}

func NewUserService(userRepository interf.UserRepository) services.UserService {
	return &UserService{UserRepository: userRepository}
}

func (u *UserService) GetUsers() ([]user.Entity, error) {
	return u.UserRepository.GetUsers()
}

func (u *UserService) CreateUser(input *user.InputResponse) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	input.Password = string(hashedPassword)
	return u.UserRepository.CreateUser(input)
}

func (u *UserService) GetUserById(id string) (user.Entity, error) {
	return u.UserRepository.GetUserById(id)
}

func (u *UserService) UpdateUser(id string, input *user.InputResponse) (int, error) {
	return u.UserRepository.UpdateUser(id, input)
}

func (u *UserService) DeleteUser(id string) (int, error) {
	return u.UserRepository.DeleteUser(id)
}

func (u *UserService) SearchUser(queryType, query string) ([]user.Entity, error) {
	return u.UserRepository.SearchUser(queryType, query)
}

func (u *UserService) Login(email, password string) (int, error) {
	return u.UserRepository.Login(email, password)
}
