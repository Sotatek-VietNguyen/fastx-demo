package services

import (
	"gorm.io/gorm"
	"user-service/src/models"
	"user-service/src/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	userRepo := repositories.NewUserRepository(db)
	return &UserService{repo: userRepo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}
