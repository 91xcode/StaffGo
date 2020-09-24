package main

import (
	"bufio"
	"code.be.staff.com/staff/StaffGo/public/common"
	"code.be.staff.com/staff/StaffGo/public/mysql"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

/*
	文件分割与入库

	先将大文件分割成N=10个小文件

	然后用10个goroutine 每个goroutine 读一个文件 将内容 写到一个chan

	然后用20个goroutine消费这个chan 然后遍历这个chan 没M条记录一起Insert数据库 数据库连接池中连接数上线设置为10


 */

const (
	FILEDIR    = "/Users/liubing/StaffGo/goroutine/Example/Files/"
	ORIGINNAME = "data.txt"
	FILE_PREFIX = "test_"
	FILE_EXTENSION = ".txt"
	LEN  = 1000

)

func main() {

	SplitBigFile()

	LoadData()

}

func SplitBigFile() {

	defer common.TimeCost()()

	var wg sync.WaitGroup

	ch := make(chan string)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go consume(&wg, i, ch)
	}

	product(ch)
	wg.Wait()
}

func product(ch chan<- string) {
	file, err := os.Open(FILEDIR + ORIGINNAME)
	defer file.Close()

	if err != nil {
		common.HandleError(err, "file open err")
	}

	//创建缓冲读取器
	reader := bufio.NewReader(file)
	for {
		//读取一行
		lineBytes, _, err := reader.ReadLine()

		if err == io.EOF {
			fmt.Println("已经读到文件末尾！")
			close(ch)
			break
		}

		lineStr := string(lineBytes)

		ch <- lineStr
	}
}

func consume(wg *sync.WaitGroup, indx int, ch <-chan string) {
	defer wg.Done()
	file, err := os.OpenFile(FILEDIR+FILE_PREFIX+strconv.Itoa(indx)+FILE_EXTENSION,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	common.HandleError(err, `os.OpenFile`)
	defer file.Close()

	totalLines := 0
	for item := range ch {
		file.WriteString(item + "\n")
		//统计写出数量
		totalLines++
		fmt.Printf("协程%+v,写入:%+v \n", indx, totalLines)
	}
}

type Info struct {
	Id int `json:"id"`
	//姓名
	Name string `json:"name"`
	//身份证号码
	Card_id string `json:"card_id"`
	//手机号
	Phone string `json:"phone"`
	//银联号
	Bank_id string `json:"bank_id"`
	//邮箱
	Email string `json:"email"`
	//地址
	Address string `json:"address"`
	//时间
	Create_at string `json:"create_at"`
}

var (
	ErrSQLLen = errors.New("insert2DB SQL len  ")
)

func initDB() {
	cfg := mysql.PoolConfig{"root:@tcp(localhost:3306)/go_demos?parseTime=true&loc=Local&charset=utf8", 10, 10}
	if err := mysql.Init(cfg); err != nil {
		log.Printf("Init(%v) error(%v)", cfg, err)
		return
	}
	sql := `create table if not exists person(
		id  int(10) unsigned NOT NULL AUTO_INCREMENT,
		name char(20),
		card_id char(18),
		phone char(11),
		bank_id char(19),
		email char(255),
		address varchar(255),
		create_at int(10),
        create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8`
	//必要时先建表
	_, err := mysql.Exec(sql)
	common.HandleError(err, "db.Exec create table")
	fmt.Println("数据表已创建")
}

func closeDB() {
	mysql.Close()
}

var (
	readingFinished int
	lock            sync.Mutex
)

func LoadData() {
	defer common.TimeCost()()

	initDB()

	var wg sync.WaitGroup
	dataCh := make(chan Info)

	/*开辟10条读取协程，分别读取不同的文件*/
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go productData(&wg, i, dataCh)
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go consumeData(&wg, i, dataCh)
	}

	wg.Wait()

	closeDB()
}

func productData(wg *sync.WaitGroup, indx int, dataCh chan<- Info) {
	defer wg.Done()

	file, err := os.Open(FILEDIR + FILE_PREFIX + strconv.Itoa(indx) + FILE_EXTENSION, )
	common.HandleError(err, `os.Open`)
	defer file.Close()

	//创建缓冲读取器
	reader := bufio.NewReader(file)
	fmt.Println("大数据文本已打开")

	for {
		//读取一行
		lineBytes, _, err := reader.ReadLine()

		if err == io.EOF {

			lock.Lock()
			defer lock.Unlock()
			readingFinished++

			fmt.Printf("readingFinished:%d\n", readingFinished)

			//读取协程全部完毕时，关闭数据管道
			if readingFinished > 9 {
				close(dataCh)
			}
			break
		}

		lineStr := string(lineBytes)

		fields := strings.Split(lineStr, ",")

		name, card_id, phone, bank_id, email, address, create_at :=
			fields[0], fields[1], fields[2], fields[3], fields[4], fields[5], fields[6]

		data := Info{Name: name, Card_id: card_id, Phone: phone, Bank_id:
			bank_id, Email: email, Address: address, Create_at: create_at}

		dataCh <- data
	}
}

func consumeData(wg *sync.WaitGroup, indx int, dataCh <-chan Info) {

	defer wg.Done()

	list := make([]*Info, 0, LEN)
	for item := range dataCh {

		list = append(list, &Info{Name: item.Name, Card_id: item.Card_id, Phone: item.Phone,
			Bank_id: item.Bank_id, Email: item.Email, Address: item.Address, Create_at: item.Create_at})
		if len(list) == LEN {

			insert2DB(list)
			cleanslice(list)
			list = make([]*Info, 0, LEN)
		}

	}

	insert2DB(list)
	cleanslice(list)
	list = make([]*Info, 0, LEN)
}

func insert2DB(list []*Info) error {
	sql := `insert into person (name,card_id,phone,bank_id,email,address,create_at) values`

	for _, item := range list {
		if item == nil {
			continue
		}

		value := `("` + item.Name + `","` + item.Card_id + `","` + item.Phone + `","` + item.Bank_id + `",
"` + item.Email + `","` + item.Address + `","` + item.Create_at + `"),`
		if (value == "") {
			continue
		}

		sql += value
	}

	if (len(sql) == 78) {
		return ErrSQLLen
	}

	//去掉最后一个逗号再加分号，形成最终SQL语句
	sql = sql[:len(sql)-1] + ";"

	//fmt.Printf("sql:%+v \n",sql)
	result, err := mysql.Exec(sql)

	if (err != nil) {
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

/*清空切片，回收内存，避免内存泄露*/
func cleanslice(list []*Info) {

	for i := 0; i < len(list); i++ {
		list[i] = nil
	}
	runtime.GC()
}
