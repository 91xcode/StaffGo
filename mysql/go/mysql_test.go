package main

import "testing"

//执行 go test -v mysql_test.go mysql.go


func TestGetAll(t *testing.T) {

	//t.Log("hello world")
	InitDB()
	result := getAll()
	// 显示用户信息
	for _, val := range result {
		t.Logf("%+v\n", val)
		t.Logf("%+v\n", val)
		//fmt.Printf("%d, %s, %s\n", val.Id, val.Username, val.Password)
	}
	CloseDB()
}