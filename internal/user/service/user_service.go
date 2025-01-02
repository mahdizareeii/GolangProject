package service

import (
	"GolangProject/internal/user"
	"GolangProject/internal/user/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers() ([]user.User, error) {
	return s.repo.FetchAllUsers()
}
