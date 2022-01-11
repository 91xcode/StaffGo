package models

type User struct {
	Name   string `xorm:"varchar(255)" json:"name"`
	Mobile string `xorm:"varchar(255)" json:"mobile"`
	Age    int64  `xorm:"int" json:"age"`
}