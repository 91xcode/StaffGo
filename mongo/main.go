package main

import (
	"os"
	"code.be.staff.com/staff/StaffGo/mongo/mongo"
	"code.be.staff.com/staff/StaffGo/public/zap"
	"code.be.staff.com/staff/StaffGo/mongo/conf"
	"time"
	bookmodel "code.be.staff.com/staff/StaffGo/mongo/model/book"
)

func Init() {
	zap.Init(&conf.Conf.Zap)
	zap.Zlogger.Info("zap init success")

	mongo.Init(&conf.Conf.Mongo)
	zap.Zlogger.Info("mongo init success")
}

func add() {

	zap.Zlogger.Infof("add start two============>")

	two := bookmodel.BookInfo{
		Title:   "标题",
		Des:     "博客描述信息2",
		Content: "博客的内容信息2",
		Img:     "https://upload-images.jianshu.io/upload_images/8679037-67456031925afca6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/700",
		Date:    time.Now(),
	}

	bookMonGoDao := mongo.NewBookMonGoDao()

	rst, err := bookMonGoDao.InsertTwo(&two);

	if err != nil {
		zap.Zlogger.Infof("err:%+v", err)
	}
	zap.Zlogger.Infof("rst:%+v", rst)

}


func getOne(){

	zap.Zlogger.Infof("getOne============>")
	//query
	query:=bookmodel.BookInfo{
		Title:"标题",
	}

	bookMonGoDao := mongo.NewBookMonGoDao()
	res, err := bookMonGoDao.GetOneByTitle(&query)
	if err != nil {
		zap.Zlogger.Infof("err:%+v", err)
	}

	zap.Zlogger.Infof("res:%+v", res)
}

func getAll(){
	zap.Zlogger.Infof("getAll============>")


	bookMonGoDao := mongo.NewBookMonGoDao()
	res, err := bookMonGoDao.GetAll()
	if err != nil {
		zap.Zlogger.Infof("err:%+v", err)
	}

	zap.Zlogger.Infof("res:%+v", res)
}


func delOne(){
	zap.Zlogger.Infof("delOne============>")

	query:=bookmodel.BookInfo{
		Title:"标题99",
	}

	bookMonGoDao := mongo.NewBookMonGoDao()
	res, err := bookMonGoDao.GetOneByTitle(&query)
	if err != nil {
		zap.Zlogger.Infof("err:%+v", err)
	}

	_id:=res.Id
	que:=bookmodel.BookInfo{
		Id:_id,
	}
	rst, err := bookMonGoDao.RemoveOne(&que)
	if err != nil {
		zap.Zlogger.Infof("err:%+v", err)
	}
	zap.Zlogger.Infof("rst:%+v", rst)
}




func BulkAdd() {

	zap.Zlogger.Infof("BulkAdd start one============>")

	params:=[]bookmodel.BookInfo{
		bookmodel.BookInfo{
			Title:   "标题99",
			Des:     "博客描述信息",
			Content: "博客的内容信息",
			Img:     "https://upload-images.jianshu.io/upload_images/8679037-67456031925afca6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/700",
			Date:    time.Now(),
		},
		bookmodel.BookInfo{
			Title:   "标题98",
			Des:     "博客描述信息",
			Content: "博客的内容信息",
			Img:     "https://upload-images.jianshu.io/upload_images/8679037-67456031925afca6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/700",
			Date:    time.Now(),
		},
	}

	bookMonGoDao := mongo.NewBookMonGoDao()
	for _, item:=range params {
		_, err := bookMonGoDao.InsertTwo(&item);
		if err != nil {
			zap.Zlogger.Infof("err:%+v", err)
		}
		zap.Zlogger.Infof("item:%+v", item)
	}
}


func FindPages(){
	bookMonGoDao := mongo.NewBookMonGoDao()

	//var resultWithPage []Data
	resultWithPage,err := bookMonGoDao.FindPageData(0, 2, nil, nil)
	if err != nil {
		zap.Zlogger.Infof("err:%+v", err)
	}
	zap.Zlogger.Infof("resultWithPage:%+v", resultWithPage)
}


func main() {

	if len(os.Args) < 2 {
		panic("params error you need  go run main.go config.ini")
	}

	if err := conf.Init(os.Args[1]); err != nil {
		panic(err)
	}

	Init()

	add()

	getOne()

	//getAll()

	//delOne()

	//delAll()

	//BulkAdd()


	//delOne()

	//FindPages()




}
