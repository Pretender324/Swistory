package db

import (
	"fmt"
	"swistory/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	dsn := "Pretender324:Riku0324#@unix(/tmp/mysql.sock)/swistory?charset=utf8"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	db.AutoMigrate(&model.Menu{})
	autoMigration()
}

// DB取得
func GetDB() *gorm.DB {
	return db
}

// マイグレーション
func autoMigration() {
	db.AutoMigrate(&model.Menu{})
}
