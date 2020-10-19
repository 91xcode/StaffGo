package main

import (
	"code.be.staff.com/staff/StaffGo/public/httpclient"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/urfave/cli"
	"encoding/json"
	"errors"
)

const api = "https://gank.io/api/v2/hot/%s/category/%s/count/%d"

var (
	ErrRes = errors.New("response error")
)

type Au struct {
	Category string `json:"category"`
	Data     []struct {
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
	} `json:"data"`
	Hot    string `json:"hot"`
	Status int    `json:"status"`
}

func main() {
	app := cli.NewApp()
	app.Name = "Girl-cli"
	app.Usage = "Girl小程序"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "hot_type, ht",
			Value: "likes",
			Usage: "可接受参数 views（浏览数） | likes（点赞数） | comments（评论数）",
		},
		cli.StringFlag{
			Name:  "category, c",
			Value: "Girl",
			Usage: "可接受参数 Article | GanHuo | Girl",
		},
		cli.IntFlag{
			Name:  "count, co",
			Value: 10,
			Usage: "可接受参数 [0,20]",
		},
	}

	app.Action = func(c *cli.Context) error {
		hot_type := c.String("hot_type")
		category := c.String("category")
		count := c.Int("count")


		uri := fmt.Sprintf(api, hot_type, category, count)

		fmt.Printf("uri:%s \n", uri)

		resp, err := httpclient.Get(uri, nil)
		if err != nil {
			fmt.Printf("err was %v", err)
			return nil
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("ErrResponse was %v", ErrRes)
			return nil
		}

		respone := Au{}
		if err := json.Unmarshal(body, &respone); err != nil {
			fmt.Printf("\nError message: %v", err)
			return nil
		}

		//fmt.Printf("respone:%+v \n", respone)


		for _,item:=range respone.Data{
			fmt.Printf("Image:%+v \n", item.Images[0])
		}

		return nil
	}
	app.Run(os.Args)
}