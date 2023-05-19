package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `gorm:"type:varchar(100);unique_index" json:"email"`
}

func (User) TableName() string {
	return "users"
}
