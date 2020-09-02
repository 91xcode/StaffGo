package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
	// "fmt"
	"code.be.staff.com/staff/StaffGo/public/zap"
	"time"
)

const (
	timeout = 60 * time.Second
	source  = "admin"
)

var (
	globalS    *mgo.Session
	Collection string
	Database   string
)

func init() {
	Collection = ""
	Database = ""
}

// MongoConfig 配置
type MongoConfig struct {
	Addr       string `json:"Addr"`
	User       string `json:"User"`
	Password   string `json:"Password"`
	DB         string `json:"DB"`
	Collection string `json:"Collection"`
	PoolSize   int    `json:"PoolSize"`
}

func Init(conf *MongoConfig) {

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{conf.Addr},
		Timeout:   timeout,
		Source:    source,
		Database:  conf.DB,
		Username:  conf.User,
		Password:  conf.Password,
		PoolLimit: conf.PoolSize,
	}
	Collection = conf.Collection
	Database = conf.DB
	zap.Zlogger.Info("init Mongo:", conf)

	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		zap.Zlogger.Fatalf("Create Mongo Session: %s\n", err)
	}
	globalS = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := globalS.Copy()
	c := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}

func getDb(db string) (*mgo.Session, *mgo.Database) {
	ms := globalS.Copy()
	return ms, ms.DB(db)
}

func IsEmpty(db, collection string) bool {
	ms, c := connect(Database, Collection)
	defer ms.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func Count(query interface{}) (int, error) {
	ms, c := connect(Database, Collection)
	defer ms.Close()
	return c.Find(query).Count()
}

func Insert(docs ...interface{}) error {
	ms, c := connect(Database, Collection)
	defer ms.Close()

	return c.Insert(docs...)
}

func FindOne(query, selector, result interface{}) error {
	ms, c := connect(Database, Collection)
	defer ms.Close()

	return c.Find(query).Select(selector).One(result)
}

func FindAll(query, selector, result interface{}) error {
	ms, c := connect(Database, Collection)
	defer ms.Close()

	return c.Find(query).Select(selector).All(result)
}

//db:操作的数据库
//collection:操作的文档(表)
//page:当前页面
//limit:每页的数量值
//query:查询条件
//selector:需要过滤的数据(projection)
//result:查询到的结果
func FindPage(page, limit int, query, selector, result interface{}) error {
	ms, c := connect(Database, Collection)
	defer ms.Close()

	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func FindIter(query interface{}) *mgo.Iter {
	ms, c := connect(Database, Collection)
	defer ms.Close()

	return c.Find(query).Iter()
}

func Update(selector, update interface{}) error {
	ms, c := connect(Database, Collection)
	defer ms.Close()

	return c.Update(selector, update)
}

func Upsert(selector, update interface{}) error {
	ms, c := connect(Database, Collection)
	defer ms.Close()

	_, err := c.Upsert(selector, update)
	return err
}

func UpdateAll(selector, update interface{}) error {
	ms, c := connect(Database, Collection)
	defer ms.Close()

	_, err := c.UpdateAll(selector, update)
	return err
}

func Remove(selector interface{}) error {
	ms, c := connect(Database, Collection)
	defer ms.Close()

	return c.Remove(selector)
}

func RemoveAll(selector interface{}) error {
	ms, c := connect(Database, Collection)
	defer ms.Close()

	_, err := c.RemoveAll(selector)
	return err
}

//insert one or multi documents
func BulkInsert(docs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(Database, Collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Insert(docs...)
	return bulk.Run()
}

func BulkRemove(selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(Database, Collection)
	defer ms.Close()

	bulk := c.Bulk()
	bulk.Remove(selector...)
	return bulk.Run()
}

func BulkRemoveAll(selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(Database, Collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.RemoveAll(selector...)
	return bulk.Run()
}

func BulkUpdate(pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(Database, Collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Update(pairs...)
	return bulk.Run()
}

func BulkUpdateAll(pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(Database, Collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.UpdateAll(pairs...)
	return bulk.Run()
}

func BulkUpsert(pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(Database, Collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Upsert(pairs...)
	return bulk.Run()
}

func PipeAll(pipeline, result interface{}, allowDiskUse bool) error {
	ms, c := connect(Database, Collection)
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.All(result)
}

func PipeOne(pipeline, result interface{}, allowDiskUse bool) error {
	ms, c := connect(Database, Collection)
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.One(result)
}

func PipeIter(pipeline interface{}, allowDiskUse bool) *mgo.Iter {
	ms, c := connect(Database, Collection)
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}

	return pipe.Iter()

}

func Explain(pipeline, result interface{}) error {
	ms, c := connect(Database, Collection)
	defer ms.Close()
	pipe := c.Pipe(pipeline)
	return pipe.Explain(result)
}
func GridFSCreate(prefix, name string) (*mgo.GridFile, error) {
	ms, d := getDb(Database)
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Create(name)
}

func GridFSFindOne(prefix string, query, result interface{}) error {
	ms, d := getDb(Database)
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Find(query).One(result)
}

func GridFSFindAll(prefix string, query, result interface{}) error {
	ms, d := getDb(Database)
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Find(query).All(result)
}

func GridFSOpen(prefix, name string) (*mgo.GridFile, error) {
	ms, d := getDb(Database)
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Open(name)
}

func GridFSRemove(prefix, name string) error {
	ms, d := getDb(Database)
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Remove(name)
}
