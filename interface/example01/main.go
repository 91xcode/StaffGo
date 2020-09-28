package main

import (
	"code.be.staff.com/staff/StaffGo/interface/example01/sdk"
	"fmt"
)


type Data struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func main(){
	cpID:= 1500
	api:=sdk.GetAPI(cpID)
	if api == nil {
		return
	}


	d:=&Data{Id:88,Name:"Hello"}

	if err := api.Get(d.Id, d.Name); err != nil {
		fmt.Printf("err:%+v",err)
	}

	fmt.Println("ok")
	return
}