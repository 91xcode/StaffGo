package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"io"
	"log"
	"os"
	"time"
)

type Meizi struct {
	Photo_href string `json:"photo_href"`
	People_href string `json:"people_href"`
}

const ResultFileName  = "./result.txt"

func run(){
	// 起始Url
	startUrl := "http://www.douban.com/photos/album/75978669/?m_start=0"
	// 创建Collector
	collector := colly.NewCollector(
		// 设置用户代理
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36"),
		colly.Async(true),  //异步
		)

	// 设置抓取频率限制
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 2 * time.Second, // 随机延迟
	})

	// 异常处理
	collector.OnError(func(response *colly.Response, err error) {
		log.Println(err.Error())
	})

	collector.OnRequest(func(request *colly.Request) {
		log.Println("start visit: ", request.URL.String())
	})



	// 解析列表
	collector.OnHTML(".article", func(element *colly.HTMLElement) {
		// 依次遍历所有的li节点
		var lists []Meizi
		element.DOM.Find(".photo_wrap").Each(func(i int, selection *goquery.Selection) {
			href, found := selection.Find(".photolst_photo").Attr("href")
			// 如果找到了详情页，则继续下一步的处理
			if found {
				ret,_:=parseDetail(collector, href)
				lists=append(lists,ret)
			}
		})
		//log.Printf("lists:%+v \n",lists)

		for _,item:=range lists{

			content := fmt.Sprintf("%s,%s\n", item.Photo_href, item.People_href)
			writeFile(ResultFileName,content)
			log.Printf("photo_href:%+v,people_href:%+v \n",item.Photo_href,item.People_href)
		}


	})


	// 查找下一页
	collector.OnHTML("div.paginator > span.next", func(element *colly.HTMLElement) {
		href, found := element.DOM.Find("a").Attr("href")
		// 如果有下一页，则继续访问
		if found {
			element.Request.Visit(element.Request.AbsoluteURL(href))
		}
	})

	// 起始入口
	collector.Visit(startUrl)

	collector.Wait()  //异步
}



/**
 * 处理详情页
 */
func parseDetail(collector *colly.Collector, url string)(rst Meizi,err error){
	collector = collector.Clone()

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 2 * time.Second,
	})

	collector.OnRequest(func(request *colly.Request) {
		log.Println("start detail visit: ", request.URL.String())
	})


	// 解析详情页数据
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		photo_href, _ := element.DOM.Find("div.image-show").Find("div img ").Attr("src")
		people_href := element.DOM.Find("div.photo_descri").Find("div.pl").Text()
		rst = Meizi{
			Photo_href:    photo_href,
			People_href:  people_href,
		}

		//fmt.Printf("%+v", rst)

	})

	collector.Visit(url)
	collector.Wait() //异步
	return rst,err
}




func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}


func writeFile(name,content string){
	/*
		 O_RDONLY int = syscall.O_RDONLY // 只读打开文件和os.Open()同义
		 O_WRONLY int = syscall.O_WRONLY // 只写打开文件
		 O_RDWR   int = syscall.O_RDWR   // 读写方式打开文件
		 O_APPEND int = syscall.O_APPEND // 当写的时候使用追加模式到文件末尾
		 O_CREATE int = syscall.O_CREAT  // 如果文件不存在，此案创建
		 O_EXCL   int = syscall.O_EXCL   // 和O_CREATE一起使用, 只有当文件不存在时才创建
		 O_SYNC   int = syscall.O_SYNC   // 以同步I/O方式打开文件，直接写入硬盘.
		 O_TRUNC  int = syscall.O_TRUNC  // 如果可以的话，当打开文件时先清空文件
	 */
	fileObj,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	if  _,err := io.WriteString(fileObj,content);err == nil {
		fmt.Println("写入文件成功:",content)
	}
}

func main(){
	run()
}