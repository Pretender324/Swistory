package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	GroupID     uint
	Level       string
	Aim         string
	Description string
}

type Practice struct {
	gorm.Model
	MenuID      int
	Menu        Menu
	Category    string
	System      string
	Distance    string
	Times       string
	Set         string
	OnTime      string
	Description string
}

type User struct {
	gorm.Model
	Name    string
	Sex     string
	Email   string
	GroupID uint
}

type Group struct {
	gorm.Model
	Name  string
	Users []User
	Menus []Menu
}