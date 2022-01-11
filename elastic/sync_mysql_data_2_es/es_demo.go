package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"reflect"
	"time"
)

var (
	es_url = "http://127.0.0.1:9200"
	es_ctx   = context.Background()
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 链接服务器
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(es_url),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetMaxRetries(5),
	)
	if err != nil {
		panic(err.Error())
	}

	// 获取基本信息
	info, code, err := client.Ping(es_url).Do(es_ctx)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(">>es-info>>", info.Version.Number, code)

	client.Delete().Index("user")

	// 是否存在索引
	exists, err := client.IndexExists("user").Do(es_ctx)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(">>index-exists>>", exists)

	// 创建索引
	if !exists {
		rst, err := client.CreateIndex("user").Do(es_ctx)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(">>index-index>>", rst.Index)
	}

	// 批量添加
	for i := 1; i <= 3; i++ {
		id := fmt.Sprintf("%d", i)
		rsp, err := client.Index().
			Index("user").
			Type(type_name).
			Id(id).
			BodyJson(&User{Name: fmt.Sprintf("name%d", i), Age: i}).
			Do(es_ctx)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(">>doc-create>>", "id:", rsp.Id, "index:", rsp.Index, "type:", rsp.Type, "seqno:",
			rsp.SeqNo, "result:", rsp.Result, "status:", rsp.Status)
	}

	// 修改
	rsp, err := client.Update().Index("user").Id("1").Doc(&User{Name: "name100", Age: 100}).Do(es_ctx)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(">>doc-update>>", rsp)

	// 删除
	rsp2, err2 := client.Delete().Index("user").Id("2").Do(es_ctx)
	if err2 != nil {
		panic(err2.Error())
	}
	fmt.Println(">>doc-delete>>", rsp2)

	// 查询所有
	rst, err := client.Search().Index("user").Do(es_ctx)
	if err != nil {
		panic(err.Error())
	}
	for _, u := range rst.Each(reflect.TypeOf(User{})) {
		fmt.Println(">>search-all>>", u.(User).Name, u.(User).Age)
	}

	// 查询条件 - 相等
	q := elastic.NewQueryStringQuery("name:name1")
	rst, err = client.Search().Index("user").Query(q).Do(es_ctx)
	if err != nil {
		panic(err.Error())
	}
	for _, u := range rst.Each(reflect.TypeOf(User{})) {
		fmt.Println(">>search-eq>>", u.(User).Name, u.(User).Age)
	}

	// // 查询条件 - 大于 - 分页
	q2 := elastic.NewBoolQuery()
	q2.Must(elastic.NewRangeQuery("age").Gt(3))
	rst, err = client.Search().
		Index("user").
		Size(2).
		From(1).
		Sort("age", false).
		Query(q2).
		Do(es_ctx)
	if err != nil {
		panic(err.Error())
	}
	for _, u := range rst.Each(reflect.TypeOf(User{})) {
		fmt.Println(">>search-eq2>>", u.(User).Name, u.(User).Age)
	}
}
