package main

import (
	"fmt"
	"code.be.staff.com/staff/StaffGo/elastic"
	"time"

	"github.com/olivere/elastic"
)

// Tweet is a structure used for serializing/deserializing data in Elasticsearch.
type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []int                 `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
	SS       float64               `json:"ss"`
	WW       []byte                `json:"ww"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"tweet":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"image":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`

func main() {
	es := myelastic.OnInitES("http://192.168.0.152:9200/")
	if es.Err != nil {
		mylog.Error(es.Err)
		panic(es.Err)
	}

	b := es.CreateIndex("xxj", mapping)
	fmt.Println(b)
	var ss Tweet
	ss.Created = time.Now()
	ss.Location = "ssss"
	ss.SS = 123.11243
	ss.WW = []byte("xxjwxc")
	ss.Tags = []int{12, 22}

	b = es.Add("wwww5", "wxc5", "", ss)
	if !b {
		fmt.Println(es.Err)
	}

	//	var ws []map[string]interface{}
	//	b = es.SearchMap("", "", `{"query":{"match_all":{}}}`, &ws)
	//	for _, v := range ws {
	//		fmt.Println(v)
	//	}
	//	fmt.Println(len(ws))

	var ws []Tweet
	b = es.Search("wwww5", "wxc5", `{"query":{"match_all":{}}}`, &ws)

	for _, v := range ws {
		fmt.Println(v)
	}

	fmt.Println(len(ws))

}
