package services

import (
	"errors"
	"github.com/YJU-OKURA/project_minori-gin-deployment-repo/models"
	"github.com/YJU-OKURA/project_minori-gin-deployment-repo/repositories"
)

const ErrUserNotFound = "user not found"

type UserService interface {
	GetApplyingClasses(userID uint) ([]models.ClassUser, error)
}

type userServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewCreateUserService(userRepo repositories.UserRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (s *userServiceImpl) GetApplyingClasses(userID uint) ([]models.ClassUser, error) {
	exists, err := s.userRepo.UserExists(userID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New(ErrUserNotFound)
	}

	return s.userRepo.GetApplyingClasses(userID)
}
