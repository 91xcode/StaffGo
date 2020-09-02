package mongo

import (
	bookmodel "code.be.staff.com/staff/StaffGo/mongo/model/book"
	"gopkg.in/mgo.v2/bson"
)




type BookMonGoDao struct {

}

func NewBookMonGoDao() *BookMonGoDao {
	return &BookMonGoDao{

	}
}




func (dao *BookMonGoDao) InsertTwo(data *bookmodel.BookInfo) (rst bool, err error) {
	data.Id = bson.NewObjectId()
	err = Insert(data)
	if err != nil {
		return false, err
	}
	return true, err
}


func (dao *BookMonGoDao) GetOneByTitle(data *bookmodel.BookInfo) (result bookmodel.BookInfo, err error)  {
	// find one with all fields
	err = FindOne(bson.M{"title": data.Title}, nil, &result)
	return
}


func (dao *BookMonGoDao) GetAll() (result bookmodel.BookInfo, err error)  {
	// find one with all fields
	err = FindAll(nil, nil, &result)
	return
}

func (dao *BookMonGoDao) RemoveOne(data *bookmodel.BookInfo)(rst bool, err error)   {
	// find one with all fields
	err = Remove(bson.M{"_id": data.Id})
	if err != nil {
		return false, err
	}
	return true, err
}

func (dao *BookMonGoDao) FindPageData(page,limit int,query, selector  interface{})(result []bookmodel.BookInfo,err error){
	err = FindPage( page, limit, query, selector, &result)
	return
}






