package main

import (
	"code.be.staff.com/staff/StaffGo/thrift/go/thrift/user"
	"github.com/apache/thrift/lib/go/thrift"
	"fmt"
	"golang.org/x/net/context"
)

const (
	NetworkAddr = "127.0.0.1:9090" //监听地址&端口
)

var ctx = context.Background()

func GetClient() *user.UserClient {
	addr := NetworkAddr
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
	}

	//protocol
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	//no buffered
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("error running client:", err)
	}

	if err := transport.Open(); err != nil {
		fmt.Println("error running client:", err)
	}


	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := user.NewUserClient(thrift.NewTStandardClient(iprot, oprot))
	return client
}

func main() {

	// 调用rpc服务提供的方法
	client := GetClient()
	rep, err := client.GetUser(ctx, 100)
	if err != nil {
		fmt.Printf("thrift err: %v\n", err)
	} else {
		fmt.Printf("Recevied: %v\n", rep)
	}
}



//func one() {
//	var transport thrift.TTransport
//	var err error
//
//	// 传输方式（需要与服务端一致）
//	transport, err = thrift.NewTSocket("localhost:9090")
//	if err != nil {
//		fmt.Printf("open socket err, %s\n", err)
//		return
//	}
//
//	// 传输协议（需要与服务端一致）
//	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
//	iProtocol := protocolFactory.GetProtocol(transport)
//	oProtocol := protocolFactory.GetProtocol(transport)
//	tClient := thrift.NewTStandardClient(iProtocol, oProtocol)
//
//	// 实际业务
//	userClient := user.NewUserClient(tClient)
//	if err := transport.Open(); err != nil {
//		fmt.Printf("Error opening socket to :9092 %s", err)
//		return
//	}
//	defer transport.Close()
//
//	// 调用rpc服务提供的方法
//	res, err := userClient.GetUser(ctx,100)
//	if err != nil {
//		fmt.Printf("get user err, %s", err)
//		return
//	}
//	fmt.Printf("result %s", res.ID)
//}