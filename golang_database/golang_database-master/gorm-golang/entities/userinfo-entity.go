package entities

import (
	"time"
)

var count uint64 = 1

// User
type User struct {
	ID         uint64    `gorm:"primary_key"`
	Username   string    `gorm:"not null"`
	Password   string    `gorm:"not null"`
	SignUpDate time.Time `gorm:"not null"`
}

// NewUser returns a new user with a new uid
func NewUser(username, password string) *User {
	u := User{
		ID:         count,
		Username:   username,
		Password:   password,
		SignUpDate: time.Now(),
	}
	count++
	return &u
}
