package spent

import (
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