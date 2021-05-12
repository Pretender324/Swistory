package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Level       string
	Aim         string
	Description string
}
