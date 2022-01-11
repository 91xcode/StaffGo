package main


import (
	"context"
	"github.com/olivere/elastic"
	"time"
)

var (
	esUrl  = "http://localhost:9200"
	ctx    = context.Background()
	client *elastic.Client
)

func init() {
	var err error

	client, err = elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(esUrl),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetMaxRetries(5),
	)

	if err != nil {
		panic(err.Error())
	}
}
