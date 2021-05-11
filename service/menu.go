package service

import (
	"swistory/model"
)

type MenuService struct{}

func (MenuService) SetMenu(menu *model.Menu) error {
	_, err := DbEngine.Insert(menu)
	if err != nil {
		return err
	}
	return nil
}

func (MenuService) GetMenuList() []model.Menu {
	tests := make([]model.Menu, 0)
	err := DbEngine.Distinct("id", "level", "aim", "description").Limit(10, 0).Find(&tests)
	if err != nil {
		panic(err)
	}
	return tests
}

func (MenuService) UpdateMenu(newMenu *model.Menu) error {
	_, err := DbEngine.Id(newMenu.Id).Update(newMenu)
	if err != nil {
		return err
	}
	return nil
}

func (MenuService) DeleteMenu(id int) error {
	menu := new(model.Menu)
	_, err := DbEngine.Id(id).Delete(menu)
	if err != nil {
		return err
	}
	return nil
}
