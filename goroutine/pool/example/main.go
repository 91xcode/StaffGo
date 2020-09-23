package main

//1千万数据 500个goroutine消费 10个数据库链接 1000条批量插入  花费时间 2m50s

//初始化slice的时候 设置一个初始的容量 效率会高很多

import (
	"code.be.staff.com/staff/StaffGo/public/mysql"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"
)
//任务对象
type task struct {
	Production
	Consumer
}
//设置消费者数目，也就是work pool大小
func(t *task)setConsumerPoolSize(poolSize int){
	t.Production.Jobs = make(chan *Job,poolSize * 10)
	t.Consumer.WorkPoolNum = poolSize
}

//任务数据对象
type Job struct {
	Data string
	Id int64
}

func NewTask(handler func(jobs chan *Job)(b bool))(t *task){
	t = &task{
		Production:Production{Jobs: make(chan *Job,100)},
		Consumer:Consumer{WorkPoolNum:100,Handler:handler},
	}
	return
}


type Production struct {
	Jobs chan *Job
}

func (c Production)AddData(data *Job){
	c.Jobs <- data
}

type Consumer struct {
	WorkPoolNum int
	Handler func(chan *Job)(b bool)
	Wg sync.WaitGroup
}

//异步开启多个work去处理任务，但是所有work执行完毕才会退出程序
func (c Consumer)disposeData(data chan *Job){
	for i:=0;i<=c.WorkPoolNum;i++{
		c.Wg.Add(1)
		go func() {
			defer func() {
				c.Wg.Done()
			}()
			c.Handler(data)

		}()
	}

	c.Wg.Wait()
}


func initDB(){
	cfg := mysql.PoolConfig{"root:@tcp(localhost:3306)/go_demos?parseTime=true&loc=Local&charset=utf8", 10, 10}
	if err := mysql.Init(cfg); err != nil {
		log.Printf("Init(%v) error(%v)", cfg, err)
		return
	}



	sql := 	`create table if not exists kfperson(
		  id int(10) unsigned NOT NULL AUTO_INCREMENT,
		  note varchar(255) NOT NULL DEFAULT '',
		  create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  num int(10) NOT NULL,
	      PRIMARY KEY (id)
    ) ENGINE=InnoDB  DEFAULT CHARSET=utf8`
	//必要时先建表
	_, err := mysql.Exec(sql)
	HandleError(err, "db.Exec create table")
	fmt.Println("数据表已创建")


}


func closeDB(){
	mysql.Close()
}


/*处理错误*/
func HandleError(err error, when string) {
	if err != nil {
		log.Fatal(err, when)
	}
}

const ThresholdNum  = 1000

func main(){

	defer timeCost()()		//注意，是对 timeCost() 返回的函数进行延迟调用，因此需要加两对小括号

	initDB()

	//1.先实现一个用于处理数据的闭包，在这里面实现自己业务
	consumerHandler := func(jobs chan *Job)(b bool) {

		//创建kfs切片，长度达到阈值时，做一次数据库写入操作
		kfs := make([]*Job, 0,ThresholdNum)

		for item := range jobs {
			//fmt.Println(item)
			////向切片中添加
			kfs = append(kfs, &Job{Data: item.Data, Id: item.Id})

			//切片中的数据量每达到1000（或者管道已关闭），就执行一次数数据库写入操作
			if len(kfs) == ThresholdNum {
				fmt.Println("this \n")
				//执行数据库插入
				insertData2DB(kfs)

				//清空切片并重新创建
				CleanSlice(kfs)
				kfs = make([]*Job, 0,ThresholdNum)
			}
		}


		//执行数据库插入
		insertData2DB(kfs)

		//清空切片并重新创建
		CleanSlice(kfs)
		kfs = make([]*Job, 0,ThresholdNum)
		return



		//单条插入
		//for job := range jobs {
		//	fmt.Println(job)
		//	//query := "insert into goworkmore (note,num) VALUE (?,?)"
		//	//rst,err := mysql.Exec(query,job.Data,job.Id)
		//	//if err!=nil{
		//	//	fmt.Printf("err:%+v \n",err)
		//	//}
		//	//fmt.Printf("rst:%+v \n",rst)
		//}
		//return
	}

	//2.new一个任务处理对象出来
	t :=NewTask(consumerHandler)
	t.setConsumerPoolSize(500)//500个协程同时消费
	//3.根据自己的业务去生产数据通过AddData方法去添加数据到生产channel,这里是1000万条数据
	go func(){
		for i := 1; i <= 10000000; i++ {
			job := new(Job)
			iStr := strconv.Itoa(i)
			job.Data = "这里面去定义你的任务数据格式"+ iStr
			id64,_:=strconv.ParseInt(iStr, 10, 64)
			job.Id = id64
			t.AddData(job)
		}
		//数据添加完毕之后 需要关闭掉channel 不然执行到最后会死锁
		close(t.Production.Jobs)
	}()

	//4.消费者消费数据
	t.Consumer.disposeData(t.Production.Jobs)

	closeDB()
}


//@brief：耗时统计函数
func timeCost() func() {
	start := time.Now()
	return func() {
		tc:=time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}


/*清空切片，回收内存，避免内存泄露*/
func CleanSlice(s []*Job) {
	//fmt.Printf("=====================len:%+v\n",len(s))
	for i := 0; i < len(s); i++ {
		s[i] = nil
	}
	runtime.GC()
}



/*将切片中的数据一次性插入DB中*/
func insertData2DB(kps []*Job) error {
	/*文本大数据中含有各种各样不合法的脏数据，做好异常的处理*/
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("!!!!!!!!!!!!!!!!!!!!", err, "!!!!!!!!!!!!!!!!!!!!")
	//	}
	//}()

	//构建SQL语句
	sqlStr := `insert into  goworkmore (note,num) values`

	/*拼接每个开房者信息到SQL语句中*/
	for _, kp := range kps {

 		if(kp == nil){
			continue
		}

		stringId := strconv.FormatInt(kp.Id,10)
		/*拼接名字为SQL语句*/
		personValue := `("` + kp.Data + `","` + stringId + `"),`
		if personValue == "" {
			continue
		}
		sqlStr += personValue
	}

	if len(sqlStr) == 41 {
		return nil
	}
	//去掉最后一个逗号再加分号，形成最终SQL语句
	sqlStr = sqlStr[:len(sqlStr)-1] + ";"

	//fmt.Printf("sqlStr:%+v",sqlStr)
	//执行SQL语句
	result, err := mysql.Exec(sqlStr)


	if err != nil {
		fmt.Println(" stmt.Exec", err)
	}

	//打印受影响的行数
	affected, err := result.RowsAffected()
	if err == nil {
		fmt.Println("执行成功，affected=", affected, err)
	} else {
		fmt.Println("!!!!!!!!!执行失败!!!!!!!!!!", err)
	}

	return nil
}





