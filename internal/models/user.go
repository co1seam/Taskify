package models

import "time"

type User struct {
	ID			uint 		`json:"-"`
	Name 		string 		`json:"name" binding:"required"`
	Username 	string 		`json:"username" binding:"required"`
	Password 	string 		`json:"password" binding:"required"`
	Email 		string 		`json:"email" binding:"required"`
	Birthdate 	string 		`json:"birthdate" binding:"required"`
	CreatedAt	time.Time	`json:"created_at" binding:"required"`
	UpdatedAt	time.Time	`json:"updated_at" binding:"required"`
}