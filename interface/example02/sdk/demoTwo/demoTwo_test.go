package demoTwo

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"fmt"
)



func TestMain(m *testing.M) {
	m.Run()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("One", testOne)
}

func testOne(t *testing.T)  {

	bd := NewTwoManagerImpl()
	list,err:=bd.GetList(2000,1,100)
	if err != nil {
		// %v 是以默认方式打印此值
		t.Errorf("Error : %v", err)
	}

	//t.Logf("list: %v", list)

	for _,item:=range list.Data {
		t.Logf("download: %v", item.URL)
		download(item.URL)
		//fmt.Printf("item.URL:%+v,item.Title:%+v \n",item.URL,item.Title)
	}

	t.Logf("finish")

}


func download(url string) (n int64, err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	}
	fmt.Println(name)
	out, err := os.Create(name+".jpg")
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return

}
