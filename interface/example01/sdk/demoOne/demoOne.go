package demoOne

import (
	"fmt"
)


type One struct {

}

func NewOne() *One{
	return &One{}
}



func (this *One)Get(id int,name string) error{

	fmt.Printf("demoOne id:%s,name:%s",id,name)

	return nil
}




