package main

import (
    "fmt"
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




func main(){
    //1.先实现一个用于处理数据的闭包，在这里面实现自己业务
    consumerHandler := func(jobs chan *Job)(b bool) {
        for job := range jobs {
            fmt.Println(job)
            time.Sleep(time.Second*1)
        }
        return
    }

    //2.new一个任务处理对象出来
    t :=NewTask(consumerHandler)
    t.setConsumerPoolSize(500)//500个协程同时消费
    //3.根据自己的业务去生产数据通过AddData方法去添加数据到生产channel,这里是1000万条数据
    go func(){
        for i := 0; i < 10000000; i++ {
            job := new(Job)
            iStr := strconv.Itoa(i)
            job.Data = "这里面去定义你的任务数据格式"+ iStr
            t.AddData(job)
        }
        //数据添加完毕之后 需要关闭掉channel 不然执行到最后会死锁
        close(t.Production.Jobs)
    }()

    //4.消费者消费数据
    t.Consumer.disposeData(t.Production.Jobs)
}



