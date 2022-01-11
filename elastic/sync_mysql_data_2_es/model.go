package main


import (
	"encoding/json"
	"time"
)

type PersonInfo struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Card_id     string    `json:"card_id"`
	Phone       string    `json:"phone"`
	Bank_id     string    `json:"bank_id"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	Create_at   int64     `json:"create_at"`
	Create_time time.Time `json:"create_time"`
}

func (*PersonInfo) TableName() string {
	return "person"
}

func (sa *PersonInfo) String() string {
	buf, _ := json.Marshal(sa)
	return string(buf)
}
