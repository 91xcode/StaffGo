package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"code.be.staff.com/staff/StaffGo/thrift/go/thrift/user"
	"github.com/apache/thrift/lib/go/thrift"
)


const (
	NetworkAddr = "127.0.0.1:9090" //监听地址&端口
)

// 用户服务
type UserService struct {
}
// 实际获取用户业务
func (u *UserService) GetUser(ctx context.Context, id int64) (r *user.UserInfo, err error) {
	fmt.Printf("user id %d\n", id)
	return &user.UserInfo{
		ID: id,
	}, nil
}
// 实际获取用户名业务
func (u *UserService) GetName(ctx context.Context, id int64) (r string, err error) {
	return "go server", nil
}

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}


func main() {

	flag.Usage = Usage
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", NetworkAddr, "Address to listen to")

	flag.Parse()

	//protocol
	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		Usage()
		os.Exit(1)
	}

	//buffered
	var transportFactory thrift.TTransportFactory
	if *buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	//framed
	if *framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	//handler
	handler := &UserService{}

	//transport,no secure
	var err error
	var transport thrift.TServerTransport
	transport, err = thrift.NewTServerSocket(*addr)
	if err != nil {
		fmt.Println("error running server:", err)
	}

	processor := user.NewUserProcessor(handler)

	fmt.Println("Starting the simple server... on ", *addr)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	err = server.Serve()

	if err != nil {
		fmt.Println("error running server:", err)
	}

}


//func one(){
//	// 服务处理（实际业务实现如上面）
//	handler := &UserService{}
//	processor := user.NewUserProcessor(handler)
//
//	// 网络传输方式（客户端与服务端需一致）
//	transport, err := thrift.NewTServerSocket("localhost:9090")
//	if err != nil {
//		fmt.Printf("thrift NewTServerSocket error, %s\n", err)
//		return
//	}
//
//	// 服务模型（采用TBinaryProtocol传输格式，客户端与服务端需一致）
//	server := thrift.NewTSimpleServer2(processor, transport)
//
//	fmt.Printf("server run ...\n")
//	server.Serve()
//}
