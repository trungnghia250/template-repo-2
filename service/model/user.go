package model

import "time"

type User struct {
	ID          string    `gorm:"column:id;primary_key"`
	UserName    string    `gorm:"column:user_name"`
	Email       string    `gorm:"column:email"`
	Phone       string    `gorm:"column:phone"`
	DateOfBirth time.Time `gorm:"column:date_of_birth"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
