package services

import (
	"github.com/google/uuid"
	"url-shotener-api/internal/models"
	"url-shotener-api/internal/repositories"
	"url-shotener-api/pkg/utils"
)

type AuthService struct {
	repository *repositories.Repository
}

func NewAuthService(repo *repositories.Repository) *AuthService {
	return &AuthService{repository: repo}
}

func (s *AuthService) Register(user models.UserInput) (uint, error) {
	newUserId := uuid.New().ID()
	return s.repository.Auth.Create(uint(newUserId), user)
}

func (s *AuthService) Login(user models.UserInput) (string, error) {
	userExist, err := s.repository.GetByUsername(user.Username)
	if err != nil {
		return "", BadUsernameOrPassword
	}
	if userExist.Password != user.Password {
		return "", BadUsernameOrPassword
	}
	return utils.GenerateJWTToken(userExist.Id)
}
