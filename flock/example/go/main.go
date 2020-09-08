package main

import (
	"fmt"
	"github.com/gofrs/flock"
	"time"
)

func main() {
	LockFile := "/Users/liubing/StaffGo/flock/example/go/a.lock"


	fileLock := flock.New(LockFile)

	locked, err := fileLock.TryLock()

	if err != nil {
		// handle locking error
		fmt.Printf("err :%s",err)
	}

	//if locked {
	//	// do work
	//	fmt.Printf("hello, world\n")
	//	time.Sleep(time.Duration(20)*time.Second)
	//	fmt.Printf("sleep over\n")
	//
	//	fileLock.Unlock()
	//}



	if locked {
		fmt.Printf("path: %s; locked: %v\n", fileLock.Path(), fileLock.Locked())

		fmt.Printf("hello, world\n")
		time.Sleep(time.Duration(20)*time.Second)
		fmt.Printf("sleep over\n")

		if err := fileLock.Unlock(); err != nil {
			// handle unlock error
			fmt.Printf("err :%s",err)
		}
	}

	fmt.Printf("path: %s; locked: %v\n", fileLock.Path(), fileLock.Locked())



}



