```
Thrift go

1.在一个 .thrift 文件内定义服务，并用 thrift 工具生成服务接口。

mkdir -p staff_go/thrift/go && cd staff_go/thrift/go

cat user.thrift
# 命名空间，可以不写也可以写多个，按照使用服务端、客户端语言来写即可
namespace go thrift.user
namespace php thrift.user

# 定义一个用户结构
struct UserInfo {
    #序号:字段类型 字段名
    1:i64 id
    2:string username
    3:string password
    4:string email
}

# 定义一个用户服务
service User{
    # 定义一个GetUser方法（接收一个用户id，返回上面定义的用户信息）
    UserInfo GetUser(1:i64 id)
    # 定义一个GetName方法（接收一个用户id，返回用户名称）
    string GetName(1:i64 id)

   # 方法定义格式：
   # 返回的类型 方法名(序号:参数类型 参数名 ... )
   # bool Test(1:i32 id, 2:string name, 3:i32 age ... )
}

2.thrift工具生成服务接口
首先
# 1.安装thrift工具。
# 我这里是Mac系统，采用 homebrew安装（这里为0.11.0版本）
➜  ~ brew install thrift

# 检测thrift是否安装成功
➜  ~ thrift -version
Thrift version 0.13.0

thrift可以帮助我们根据.thrift文件生产不同语言服务接口，就像下面这样：
thrift -out . --gen go user.thrift

    -out . 指定输出目录为当前目录
    --gen go user.thrift 指定使用go生成器，根据user.thrift文件，生成go语言代码  如下图
 tree .
.
├── README.md
├── thrift
│   └── user
│       ├── GoUnusedProtection__.go
│       ├── user-consts.go
│       ├── user-remote
│       │   └── user-remote.go
│       └── user.go
└── user.thrift

thrift文件是rpc开发的第一步，第二步按照服务接口实现服务端。

3.Thrift go服务端
注意 这里使用thrift(0.13.0)开发， 所以这个go包也必须是这个版本

export GO111MODULE=on  
go mod int 
mkdir -p server client


3.1 根据服务，查看需要实现接口（我这里只需要实现两个方法）

3.2 实现服务接口（go语言隐式实现👍）
这里定义一个用户服务，实现这两个方法即实现这个接口

cat server/server.go

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"staff_go/thrift/go/thrift/user"
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

第一步实现需要定义的服务接口，第二步启动一个thrift server。服务端至此开发完毕

4.Thrift go客户端
开发客户端代码

⚠️传输方式、协议需要与服务端一致。

 cat server/server.go
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"staff_go/thrift/go/thrift/user"
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
➜  thrift git:(master) ✗ cat client/client.go 
package main

import (
	"staff_go/thrift/thrift/user"
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

5.测试rpc客户端
开启2个终端
go run server/server.go 
 
go run client/client.go

运行客户端，正常返回，至此客户端开发完毕
```