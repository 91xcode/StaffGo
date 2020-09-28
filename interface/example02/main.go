package main

import (
	"code.be.staff.com/staff/StaffGo/interface/example02/sdk"
	"fmt"
)


type Data struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func main(){
	cpID:= 1500
	number:=5000


	//cpID:= 3000
	//number:=2000
	api:=sdk.GetAPI(cpID)
	if api == nil {
		return
	}

	list,err := api.GetList(number,1,10)

	if err != nil {
		fmt.Printf("err:%+v",err)
	}

	fmt.Printf("list:%+v",list)
	return
}