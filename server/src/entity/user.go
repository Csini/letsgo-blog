package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int64  //only because of gorm
	Username string `gorm:"primaryKey;unique;not null;type:varchar(15)"`
	Pw       string
}
