package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type UserDetails struct {
	gorm.Model
	// ID        int
	FirstName string
	LastName  string
	Mobile    string
	Email     string
	// Created_on gorm.SoftDeleteUpdateClause
}

func (u *UserDetails) SaveUser() (*UserDetails, error) {
	err := DB.Create(&u).Error

	if err != nil {
		return &UserDetails{}, err
	}

	return u, nil
}
