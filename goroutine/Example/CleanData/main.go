package main

import (
	"bufio"
	"code.be.staff.com/staff/StaffGo/public/common"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

const  (
	FILEDIR = "/Users/liubing/StaffGo/goroutine/Example/Files/"
	ORIGINNAME = "data.txt"
	FILE_EXTENSION = ".txt"
)


/*

数据读入、清洗、分类

主协程负责逐行读取文本大数据
另外开设34条子协程，负责对不同省份文件进行写入，从34个不同的管道中扫描数据并写出文件；
主协程根据身份证号反映的不同省份，将读入的信息丢入不同的管道，由对应的子协程进行文件写出；
当文件读取完毕时，关闭所有的数据管道（通知子协程停止数据扫描）；
主协程通过等待组等待所有子协程完成任务；
 */



/*省份结构体*/
type Province struct {
	//省的Id（身份证号的前两位）
	Id string

	//省份名称
	Name string

	//省份对应的文件
	File *os.File

	//写入该文件的数据管道
	chanData chan string
}

func main(){

	defer common.TimeCost()()



	//创建省份数据map，通过Id查询省信息
	pMap := make(map[string]*Province)

	/*
	为每个省创建一个文件
	为每个省创建一个Province实例，丢入map
	*/
	ps := []string{"北京市11", "天津市12", "河北省13", "山西省14", "内蒙古自治区15", "辽宁省21",
		"吉林省22", "黑龙江省23", "上海市31", "江苏省32", "浙江省33", "安徽省34", "福建省35", "江西省36",
		"山东省37", "河南省41", "湖北省42", "湖南省43", "广东省44", "广西壮族自治区45", "海南省46",
		"重庆市50", "四川省51", "贵州省52", "云南省53", "西藏自治区54", "陕西省61", "甘肃省62", "青海省63",
		"宁夏回族自治区64", "新疆维吾尔自治区65", "台湾省71", "香港特别行政区81", "澳门特别行政区91",
	}
	for _, p := range ps {

		//截取省份名称和省份Id
		name := p[:len(p)-2]
		id := p[len(p)-2:]

		//创建省份对象，并丢入map
		province := Province{Id: id, Name: name}
		pMap[id] = &province

		//为每个省份关联一个预备写入的数据文件
		file, _ := os.OpenFile(FILEDIR+province.Name+FILE_EXTENSION,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		province.File = file
		defer file.Close()

		//为当前省份关联换一个数据管道
		province.chanData = make(chan string, 100)
	}

	wg:=sync.WaitGroup{}
	/*为每个省创建一条数据写入协程，从各自的管道中读取数据并写入对应的文件*/
	for _, province := range pMap {
		wg.Add(1)
		go consume(&wg,province)
	}

	product(pMap)

	wg.Wait()

}



func product(pMap map[string]*Province){

	file,_:=os.Open(FILEDIR+ORIGINNAME)
	defer file.Close()

	//创建缓冲读取器
	reader := bufio.NewReader(file)
	for{
		//读取一行
		lineBytes, _, err := reader.ReadLine()

		//读取完毕时，关闭所有数据管道，并退出读取
		if err == io.EOF {
			fmt.Println("已经读到文件末尾！")

			/*遍历关闭所有数据管道（以通知子协程停止数据扫描）*/
			for _, province := range pMap {
				close(province.chanData)
				fmt.Println(province.Name, "管道已关闭", province.chanData)
			}
			break
		}


		//拿出省份ID
		lineStr := string(lineBytes)
		fieldsSlice := strings.Split(lineStr, ",")
		id := fieldsSlice[1][0:2]

		//根据Id查询得到省份，进而向其管道中写入当前行
		if province, ok := pMap[id]; ok {
			province.chanData <- (lineStr + "\n")
		} else {
			//这里其实也是不合法的数据
			fmt.Println("莫名其妙的省", id)
		}

	}
}


func consume(wg *sync.WaitGroup,province *Province){
	defer wg.Done()


	//扫描管道中的数据，管道关闭时循环结束
	for lineStr := range province.chanData {
		province.File.WriteString(lineStr)
		fmt.Print(province.Name, "写入", lineStr)
	}
	fmt.Println(province.Name, "管道遍历已结束", province.chanData)
}

