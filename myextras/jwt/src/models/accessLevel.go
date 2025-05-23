package models

import (
	"gorm.io/gorm"
)

type AccessLevel struct {
	gorm.Model
	Name        string  `json:"name" gorm:"uniqueIndex;not null"` // Display name for access level, unique
	Description string  `json:"description" gorm:"not null"`      // Description of the access level
	Active      bool    `json:"active" gorm:"default:true"`       // Active status
	Users       []*User `gorm:"many2many:user_access_levels;"`    // Many-to-many relationship with User. It is a back reference for easy querying.
}

// This is not necessary the grom tagging creates it automatically.
// type UserAccessLevel struct {
// 	Users        []User        `gorm:"many2many:user_access_levels;"`
// 	AccessLevels []AccessLevel `gorm:"many2many:user_access_levels;"` // Many-to-many relationship with AccessLevel
// }
