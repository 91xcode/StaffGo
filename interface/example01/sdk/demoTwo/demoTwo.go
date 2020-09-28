package demoTwo

import "fmt"

type Two struct {

}

func NewTwo() *Two{
	return &Two{}
}



func (this *Two)Get(id int,name string) error{

	fmt.Printf("demoTwo id:%s,name:%s",id,name)

	return nil
}