package main

//go run .

import (
	"fmt"
	"github.com/olivere/elastic"
	"sync"
)

var wg sync.WaitGroup
var MaxId = int64(0)
var pageSize = 100
var data = make(chan []PersonInfo)
var index = "person_info"

func init() {
	//client.DeleteIndex(index).Do(ctx)
	//client.CreateIndex(index).Do(ctx)
}

func main() {
	fmt.Println(">>start>>")

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			rows := getRows()
			fmt.Println(len(rows))
			if rows == nil || len(rows) == 0 {
				close(data)
				break

			}
			MaxId = rows[len(rows)-1].Id
			data <- rows
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case rows, ok := <-data:
				if !ok {
					return
				}
				el := client.Bulk().Index(index)
				for _, row := range rows {
					fmt.Printf("row:%+v \n", row)
					el.Add(elastic.NewBulkIndexRequest().Id(fmt.Sprintf("%d", row.Id)).Doc(row))
				}
				if _, err := el.Do(ctx); err != nil {
					fmt.Printf("err:%+v \n", err.Error())
				}
			}
		}
	}()

	wg.Wait()

	fmt.Println(">>done>>")
}

func getRows() []PersonInfo {
	rows := []PersonInfo{}
	err := db.Where("id>?", MaxId).Limit(pageSize).Find(&rows)
	if err != nil {
		fmt.Printf("err:%+v", err)
	}
	return rows
}
