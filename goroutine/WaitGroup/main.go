package main

//Go控制并发有三种种经典的方式，一种是通过channel通知实现并发控制 一种是WaitGroup，另外一种就是Context

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	DemoWaitGroup()
}

func DemoWaitGroup() {
	t1 := time.Now()
	var wg sync.WaitGroup
	urllist := []string{
		"www.baidu.com",
		"www.zhihu.com",
	}

	for _, item := range urllist {
		wg.Add(1)
		go worker(item, &wg)
	}

	wg.Wait()
	spend_time := time.Since(t1)

	fmt.Printf("done,spend_time:%s", spend_time)
}

func worker(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("start worker url:%s\n", url)
	time.Sleep(time.Second)
	fmt.Printf("finish worker url:%s\n", url)
	return
}



func test(){
	var wg sync.WaitGroup
	urlList:=[]string{
		"a.com","b.com",
	}

	for _,item:=range urlList{
		wg.Add(1)
		go workers(&wg,item)
	}

	wg.Wait()
}

func workers(wg *sync.WaitGroup ,item string){
	defer wg.Done()
	fmt.Printf("url:%s",item)
	return
}