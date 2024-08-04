package services

import (
	"github.com/co1seam/taskify/internal/models"
	"github.com/co1seam/taskify/internal/repository"
)

type Authorization interface{
	CreateUser(user models.User) (int, error)
}

type Subscriptions interface{

}

type Notification interface{

}

type Service struct{
	Authorization
	Subscriptions
	Notification
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		
	}
} 