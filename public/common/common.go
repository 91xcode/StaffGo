package common

import (
	"log"
	"time"
	"fmt"
)

//@brief：耗时统计函数
func TimeCost() func() {
	start := time.Now()
	return func() {
		tc:=time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}


/*处理错误*/
func HandleError(err error, when string) {
	if err != nil {
		log.Fatal(err, when)
	}
}