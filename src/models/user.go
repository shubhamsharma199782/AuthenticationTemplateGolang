package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int       `gorm:"primary key; autoIncrement" json:"id"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	Salt        string    `json:"salt"`
	PhoneNumber string    `gorm:"unique" json:"phoneNumber"`
	Coins       int       `json:"coins"`
	IsActive    bool      `gorm:"default true" json:"isActive"`
	LastLogin   time.Time `gorm:"default current_timestamp" json:"lastLogin"`
	CreatedAt   time.Time `gorm:"default current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default current_timestamp" json:"updated_at"`
}

// omiting some properties from UserResponse
type UserResponse struct {
	User
	Password  interface{} `json:"password,omitempty"`
	Salt      interface{} `json:"salt,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
	UpdatedAt interface{} `json:"updated_at,omitempty"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
