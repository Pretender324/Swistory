package service

import (
	"swistory/api/db"
	"swistory/api/model"
)

type UserService struct{}

func (UserService) SetUser(user *model.User) error {
	db := db.GetDB()
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (UserService) GetUserList() []model.User {
	tests := make([]model.User, 0)
	db := db.GetDB()
	result := db.Distinct("id", "level", "aim", "description").Limit(10).Find(&tests)
	if result.Error != nil {
		panic(result.Error)
	}
	return tests
}

func (UserService) UpdateUser(newUser *model.User) error {
	db := db.GetDB()
	result := db.First(&newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (UserService) DeleteUser(id int) error {
	db := db.GetDB()
	result := db.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
