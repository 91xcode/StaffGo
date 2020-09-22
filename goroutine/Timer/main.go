package main

//30s写文件一次

import (
	"code.be.staff.com/staff/StaffGo/public/times"
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const Interval  = 3 // 30s一次


const ResultFileName  = "a.txt"

var urlList []string

func test() (result int) {
	defer func() {
		result++
	}()
	return 1
}

func test1() (result int) {
	t := 7
	defer func() {
		t = t + 5
	}()
	return t
}

func main(){

	//fmt.Println(test())
	//fmt.Println(test1())

	fmt.Printf("start main .\n")

	urlList = []string{
		"baidu0.com",
		"zhihu0.com",
		"baidu1.com",
		"zhihu1.com",
	}

	stop:=make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go Worker(stop,&wg)

	//true 主动停止goroutine is_stop false 程序运行完毕之后自动停止goroutine
	//is_auto:= false
	//if(is_auto){
	//	close(stop)
	//}

	defer wg.Wait()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		s := <-c
		fmt.Printf("get a signal %s \n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			close(stop)
			fmt.Printf("exit \n")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}



}


func Worker(stop <-chan bool,wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.NewTimer(time.Second*time.Duration(Interval))
loop:
	for true {
		select {
		case <-stop:
			timer.Stop()
			fmt.Printf("exit worker .")
			break loop
		case <-timer.C:
			if err := doRun(); err != nil {
				fmt.Printf("doRun() err(%v) ")
			}
		}
		timer.Reset(time.Second*time.Duration(Interval))
	}
}


func doRun() error {
	for _,item:=range urlList {
		content := fmt.Sprintf("%s\n", item)
		writeFile(ResultFileName,content)
	}

	return nil
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
		fmt.Printf("time:%s,写入文件成功:%s",times.GetNormalTimeString(times.GetTime()),content)
	}
}






