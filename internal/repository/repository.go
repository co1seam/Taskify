package repository

import (
	"github.com/co1seam/taskify/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface{
	CreateUser(user models.User) (int, error)
}

type Subscriptions interface{

}

type Notification interface{

}

type Repository struct{
	Authorization
	Subscriptions
	Notification
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}