package services

import (
	"fastx-api/user-service/src/models"
	"fastx-api/user-service/src/repositories"
	"gorm.io/gorm"
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
