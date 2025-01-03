package service

import (
	"GolangProject/internal/user"
	"GolangProject/internal/user/repository"
	"GolangProject/pkg/response"
	"net/http"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers() response.BaseResponse[[]user.User] {
	userResult, errorResult := s.repo.GetUsers()
	if errorResult != nil {
		return response.ErrorResponse[[]user.User](&errorResult, http.StatusInternalServerError)
	}
	return response.SuccessResponse[[]user.User](userResult, http.StatusOK)
}

func (s *UserService) AddNewUser(shouldBindJson error, newUser *user.User) response.BaseResponse[int] {
	if shouldBindJson != nil {
		return response.ErrorResponse[int](&shouldBindJson, http.StatusBadRequest)
	}
	userIdResult, err := s.repo.AddNewUser(newUser)
	if err != nil {
		return response.ErrorResponse[int](&err, http.StatusInternalServerError)
	}

	return response.SuccessResponse[int](userIdResult, http.StatusOK)
}
