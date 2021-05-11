package model

type Menu struct {
	Id          int64  `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	Level       string `xorm:"varchar(40)" json:"level" form:"level"`
	Aim         string `xorm:"varchar(40)" json:"aim" form:"aim"`
	Description string `xorm:"varchar(40)" json:"description" form:"description"`
}
