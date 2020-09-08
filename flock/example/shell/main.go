package main
import (
	"fmt"
	"time"
)
func main() {
	fmt.Printf("hello, world\n")
	time.Sleep(time.Duration(10)*time.Second)
	fmt.Printf("sleep over\n")

}