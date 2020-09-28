package demoAks

import (
	"code.be.staff.com/staff/StaffGo/public/httpclient"
	"code.be.staff.com/staff/StaffGo/public/zap"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)


type AksManagerImpl struct {
}




func NewAksManagerImpl() *AksManagerImpl{
	return &AksManagerImpl{}
}


var (
	ErrResponse = errors.New("response error")
)

const (
	testAPI="https://gank.io/api/v2/hot/likes/category/Girl/count/%d"
)


type Datum struct {
	ID          string   `json:"_id"`
	Author      string   `json:"author"`
	Category    string   `json:"category"`
	CreatedAt   string   `json:"createdAt"`
	Desc        string   `json:"desc"`
	Images      []string `json:"images"`
	LikeCounts  int      `json:"likeCounts"`
	PublishedAt string   `json:"publishedAt"`
	Stars       int      `json:"stars"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	URL         string   `json:"url"`
	Views       int      `json:"views"`
}

type Response struct {
	Data        []Datum `json:"data"`
	Category        string   `json:"category"`
	Hot   string   `json:"hot"`
	Status      int64   `json:"status"`
}


func(v *AksManagerImpl)GetList(count int)(Response, error){


	var respone Response



	uri := fmt.Sprintf(testAPI, count)

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


	//fmt.Printf("respone:%+v \n",respone)
	//
	//for _,item:=range respone.Data {
	//	fmt.Printf("item.URL:%+v,item.Title:%+v \n",item.URL,item.Title)
	//}

	return respone,nil
}




