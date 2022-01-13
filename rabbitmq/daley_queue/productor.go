package main

// 这里的关键点就是将要延时的消息发送到过期队列当中, 然后监听的是过期队列转发到的 exchange 下的队列 正常情况就是始终监听一个队列,然后把过期消息发送到延时队列中,当消息到达时间后就把消息发到正在监听的队列
// 执行顺序 2个终端分别 go run comsumer.go  go run productor.go

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/admin")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	body := bodyFrom(os.Args)
	// 将消息发送到延时队列上
	err = ch.Publish(
		"", 				// exchange 这里为空则不选择 exchange
		"test_delay",     	// routing key
		false,  			// mandatory
		false,  			// immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Expiration: "10000",	// 设置五秒的过期时间
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [✅] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello9999"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
