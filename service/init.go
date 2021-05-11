package service

import (
	"errors"
	"fmt"
	"log"
	"swistory/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := "Pretender324:Riku0324#@unix(/tmp/mysql.sock)/swistory?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	err = DbEngine.Sync2(new(model.Menu))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("init data base ok")
}
