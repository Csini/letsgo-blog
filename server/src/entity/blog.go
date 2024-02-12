package entity

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	ID      int64  `gorm:"primaryKey;unique;not null"`
	Content string `gorm:"not null;type:varchar(500)"`
	User_ID string `gorm:"not null"`
	User    User   `gorm:"ForeignKey:User_ID"`
}
