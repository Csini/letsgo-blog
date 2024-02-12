package entity

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID      int64  `gorm:"primaryKey;unique;not null"`
	Content string `gorm:"not null;type:varchar(100)"`
	User_ID string `gorm:"not null"`
	Blog_ID int    `gorm:"not null"`
	User    User   `gorm:"ForeignKey:User_ID"`
	Blog    Blog   `gorm:"ForeignKey:Blog_ID"`
}
