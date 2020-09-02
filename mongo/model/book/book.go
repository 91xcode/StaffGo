package book

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type BookInfo struct {
	Id      bson.ObjectId `bson:"_id"`
	Title   string        `bson:"title"`
	Des     string        `bson:"des"`
	Content string        `bson:"content"`
	Img     string        `bson:"img"`
	Date    time.Time     `bson:"date"`
}
