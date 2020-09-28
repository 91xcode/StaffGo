package demoOne

import (
	"code.be.staff.com/staff/StaffGo/interface/example02/sdk/model"
	"code.be.staff.com/staff/StaffGo/public/httpclient"
	"code.be.staff.com/staff/StaffGo/public/zap"
	"errors"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)


type OneManagerImpl struct {
	Keys map[int]*Info
}


type Info struct {
	Category string
	Types string
}


const (
	NUMBER1 = 5000
	NUMBER2 = 5100
)


func NewOneManagerImpl() *OneManagerImpl{
	return &OneManagerImpl{
		Keys: map[int]*Info{
			NUMBER1: &Info{
				Category:   "GanHuo",
				Types:  "Android",
			},
			NUMBER2: &Info{
				Category:   "GanHuo",
				Types:  "iOS",
			},
		},
	}
}


var (
	ErrResponse = errors.New("response error")
)

const (
	testAPI="https://gank.io/api/v2/data/category/%s/type/%s/page/%d/count/%d"
)


func(v *OneManagerImpl)GetList(number,page,count int)(model.Response, error){


	var respone model.Response

	category := v.Keys[number].Category
	types := v.Keys[number].Types

	uri := fmt.Sprintf(testAPI, category, types,page,count)

	fmt.Printf("uri:%+v \n",uri)

	resp, err := httpclient.Get(uri, nil)
	if err != nil {
		zap.Zlogger.Info("httpclient.GetData(%v, nil) error(%v)", uri, err)
		return respone,err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zap.Zlogger.Info("ioutil.ReadAll(resp.Body) error(%v)", err)
		return respone,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		zap.Zlogger.Info("request momo return failed: uri:%s status code:%v body:%s", uri, resp.StatusCode, body)
		return respone,ErrResponse
	}


	if err := json.Unmarshal(body, &respone); err != nil {
		zap.Zlogger.Info("json.Unmarshal(%s,&resp) error(%v)", body, err)
		return respone,err
	}


	fmt.Printf("respone:%+v \n",respone)

	for _,item:=range respone.Data {
		fmt.Printf("item.URL:%+v,item.Title:%+v \n",item.URL,item.Title)
	}

	return respone,nil
}




