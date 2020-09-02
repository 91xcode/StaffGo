package main

import (
	"fmt"
	"github.com/go-ini/ini"
)


type List struct {
	AppModel string `ini:"app_model"`
	Type     int    `ini:"type"`
	Sql    Mysql    `ini:"mysql"`
}
type Mysql struct {
	Name string `ini:"name"`
	Pass string `ini:"pass"`
}

func main(){
	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfg.Section("").Key("app_model").Value())

	fmt.Println(cfg.Section("mysql").Key("name").Value())
	fmt.Println(cfg.Section("mysql").Key("pass").Value())

	//将文件中的type改为3，我们去查看一下my.ini，成功被改为3
	cfg.Section("").Key("type").SetValue("3")
	err = cfg.SaveTo("my.ini")
	if err != nil {
		fmt.Println("文件保存错误", err)
	}
	//假如我们只允许值为张三或者李四，如果用户设置的名称不在这两个里面，那么就默认为张三，代码可以如下编写

	fmt.Println(cfg.Section("mysql").Key("name").In("张三", []string{"张三", "李四"}))

	p := List{}
	err = cfg.MapTo(&p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)

}