package services

import (
	"web1/internal/core"
	"web1/internal/repositories/memory"
)

type UserService struct {
	userRepository memory.UserRepository
}

func NewUserService(userRepository memory.UserRepository) *UserService {
	userService := &UserService{userRepository: userRepository}
	return userService
}

func (userService *UserService) addUser(user *core.User) {
	userService.userRepository.AddUser(user)
}

func (userService *UserService) GetUsers() []*core.User {
	return userService.userRepository.GetAll()
}

func (userService *UserService) GetById(id int) *core.User {
	return userService.userRepository.GetById(id)
}
