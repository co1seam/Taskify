package services

import (
	"crypto/sha1"
	"fmt"

	"github.com/co1seam/Taskify/service/backend/internal/models"
	"github.com/co1seam/Taskify/service/backend/internal/repository"
)

const salt = "ejou5uethoh1Nah4piiy2Diephoh3e"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}