package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`         // Display name for user, not unique
	Password string `json:"password" gorm:"not null"`     // Password
	Email    string `json:"email" gorm:"unique;not null"` // Unique email
}

// type User struct {
// 	ID        int       `json:"id" gorm:"primaryKey"`
// 	Name      string    `json:"name" gorm:"not null"`             // Display name for user, not unique
// 	Password  string    `json:"password" gorm:"not null"`         // Password
// 	Email     string    `json:"email" gorm:"unique;not null"`     // Unique email
// 	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"` // Created at timestamp
// 	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Updated at timestamp
// 	DeletedAt time.Time `json:"deleted_at"`                       // Deleted at timestamp
// }
