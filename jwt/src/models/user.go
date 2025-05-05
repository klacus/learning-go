package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string        `json:"name" gorm:"not null"`                              // Display name for user, not unique
	Password     string        `json:"password" gorm:"not null"`                          // Password
	Email        string        `json:"email" gorm:"uniqueIndex;not null"`                 // Unique email
	Active       bool          `json:"active" gorm:"default:true"`                        // Active status
	AccessLevels []AccessLevel `json:"access_levels" gorm:"many2many:user_access_levels"` // Many-to-many relationship with AccessLevel
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
