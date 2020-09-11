package main

//Go控制并发有三种种经典的方式，一种是通过channel通知实现并发控制 一种是WaitGroup，另外一种就是Context

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	DemoChannel()
}



func DemoChannel(){

	stop:=make(chan bool)
	var wg sync.WaitGroup

	urlList := []string{
		"baidu.com",
		"zhihu.com",
	}

	for _,item:=range urlList  {
		wg.Add(1)
		go consumer(stop,item,&wg)
	}


	//true 主动停止goroutine is_stop false 程序运行完毕之后自动停止goroutine
	is_auto:= false
	if(is_auto){
		close(stop)
	}


	wg.Wait()

	fmt.Printf("done")
}

func consumer(stop <-chan bool,url string ,wg *sync.WaitGroup ){
	defer wg.Done()
	for  {
		select {
		case <-stop:
			fmt.Println("exit sub goroutine")
			return
		default:
			fmt.Printf("running goroutine..url:%s\n",url)
			time.Sleep(time.Second*20)
			return
		}
	}
}