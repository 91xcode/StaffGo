package main

//Go控制并发有三种种经典的方式，一种是通过channel通知实现并发控制 一种是WaitGroup，另外一种就是Context

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	DemoContext()
}

type favContextKey string

func DemoContext() {
	wg := &sync.WaitGroup{}
	values := []string{
		"https://www.baidu.com/",
		"https://www.zhihu.com/",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for _, url := range values {
		wg.Add(1)
		subCtx := context.WithValue(ctx, favContextKey("url"), url)
		go reqURL(subCtx, wg)
	}
	//go func() {
	//	time.Sleep(time.Second * 2)
	//	cancel()
	//}()

	//true 主动停止goroutine is_stop false 程序运行完毕之后自动停止goroutine
	is_auto:= false
	if(is_auto){
		cancel()
	}


	wg.Wait()
	fmt.Println("exit main goroutine")

}

func reqURL(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	url, _ := ctx.Value(favContextKey("url")).(string)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("stop getting url:%s\n", url)
			return
		default:
			r, err := http.Get(url)
			if r.StatusCode == http.StatusOK && err == nil {
				body, _ := ioutil.ReadAll(r.Body)
				subCtx := context.WithValue(ctx, favContextKey("resp"), fmt.Sprintf("%s%x", url, md5.Sum(body)))
				wg.Add(1)
				go showResp(subCtx, wg)
			}
			r.Body.Close()
			//启动子goroutine是为了不阻塞当前goroutine，这里在实际场景中可以去执行其他逻辑，这里为了方便直接sleep一秒
			// doSometing()
			time.Sleep(time.Second*20)
			return
		}

	}
}

func showResp(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop showing resp")
			return
		default:
			//子goroutine里一般会处理一些IO任务，如读写数据库或者rpc调用，这里为了方便直接把数据打印
			fmt.Println("printing: ", ctx.Value(favContextKey("resp")))
			time.Sleep(time.Second)
			return
		}
	}
}
