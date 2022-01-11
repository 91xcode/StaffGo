```golang grpc

服务端和客户端都是go

1. 安装protoc

brew install protobuf

2. 安装protoc-gen-go
protoc依赖该工具生成代码

go get -u github.com/golang/protobuf/protoc-gen-go
gogoprotobuf的protoc-gen-gofast插件生成的文件更复杂，性能也更高，安装如下
go get github.com/gogo/protobuf/protoc-gen-gofast



3. 安装grpc包

grpc-go包含了Go的grpc库。

go get google.golang.org/grpc 

验证是否安装好 开2个终端 cd $GOPATH/src/ 


 go run google.golang.org/grpc/examples/helloworld/greeter_server/main.go

 go run google.golang.org/grpc/examples/helloworld/greeter_client/main.go

可能会被墙掉了，使用如下方式手动安装。

git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto

cd $GOPATH/src/
go install google.golang.org/grpc





mkdir $GOPATH/src/staff_go/grpc/go/user -p && cd $GOPATH/src/staff_go/grpc/go


.
├── README.me
├── client.go
├── server.go
├── user
│   └── user.pb.go
└── user.proto


定义服务接口
vim user.proto

syntax = "proto3";
 
option go_package = "./user;user";
package user;

service User {
    rpc UserIndex(UserIndexRequest) returns (UserIndexResponse) {}
    rpc UserView(UserViewRequest) returns (UserViewResponse) {}
    rpc UserPost(UserPostRequest) returns (UserPostResponse) {}
    rpc UserDelete(UserDeleteRequest) returns (UserDeleteResponse) {}
}

message UserEntity {
    string name = 1;
    int32 age = 2;
}

message UserIndexRequest {
    int32 page = 1;
    int32 page_size = 2;
}

message UserIndexResponse {
    int32 err = 1;
    string msg = 2;
    repeated UserEntity data = 3;
}

message UserViewRequest {
    int32 uid = 1;
}

message UserViewResponse {
    int32 err = 1;
    string msg = 2;
    UserEntity data = 3;
}

message UserPostRequest {
    string name = 1;
    string password = 2;
    int32 age = 3;
}

message UserPostResponse {
    int32 err = 1;
    string msg = 2;
}

message UserDeleteRequest {
    int32 uid = 1;
}

message UserDeleteResponse {
    int32 err = 1;
    string msg = 2;
}

生成接口库
protoc-gen-go，这个工具用来将 .proto 文件转换为 Golang 代码

protoc -I. --go_out=plugins=grpc:. user.proto

编写服务端
vim server.go

package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"code.be.staff.com/staff/StaffGo/grpc/go/user"
)

const (
	port  =":7788"
	)
type UserService struct {
	// 实现 User 服务的业务对象
	user.UnimplementedUserServer
}

// UserService 实现了 User 服务接口中声明的所有方法
func (userService *UserService)UserIndex(ctx context.Context, in *user.UserIndexRequest)(*user.UserIndexResponse, error){
	log.Printf("receive user index request: page %d page_size %d", in.Page, in.PageSize)

	return &user.UserIndexResponse{
		Err: 0,
		Msg: "success",
		Data: []*user.UserEntity{
			{Name: "big_cat", Age: 28},
			{Name: "sqrt_cat", Age: 29},
		},
	}, nil
}


func (userService *UserService)UserView(ctx context.Context, in *user.UserViewRequest)(*user.UserViewResponse, error){
	log.Printf("receive user view request: uid %d", in.Uid)

	return &user.UserViewResponse{
		Err:  0,
		Msg:  "success",
		Data: &user.UserEntity{Name: "james", Age: 28},
	}, nil
}


func (userService *UserService)UserPost(ctx context.Context, in *user.UserPostRequest)(*user.UserPostResponse, error){
	log.Printf("receive user post request: name %s password %s age %d", in.Name, in.Password, in.Age)

	return &user.UserPostResponse{
		Err: 0,
		Msg: "success",
	}, nil
}


func (userService *UserService)UserDelete(ctx context.Context, in *user.UserDeleteRequest)(*user.UserDeleteResponse, error){
	log.Printf("receive user delete request: uid %d", in.Uid)

	return &user.UserDeleteResponse{
		Err: 0,
		Msg: "success",
	}, nil
}



func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 RPC 服务容器
	grpcServer := grpc.NewServer()

	// 为 User 服务注册业务实现 将 User 服务绑定到 RPC 服务容器上
	user.RegisterUserServer(grpcServer, &UserService{})

	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	//reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


编写客户端
vim client.go
package main

import (
	"google.golang.org/grpc"
	"log"
	"code.be.staff.com/staff/StaffGo/grpc/go/user"
	"time"
	"context"
	"fmt"
)

const  (
	address = "localhost:7788"
)

func main(){
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	userClient := user.NewUserClient(conn)


	// 设定请求超时时间 3s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()


	// UserIndex 请求
	userIndexReponse, err := userClient.UserIndex(ctx, &user.UserIndexRequest{Page: 1, PageSize: 12})
	if err != nil {
		log.Printf("user index could not greet: %v", err)
	}


	if 0 == userIndexReponse.Err {
		log.Printf("user index success: %s", userIndexReponse.Msg)
		// 包含 UserEntity 的数组列表
		userEntityList := userIndexReponse.Data
		for _, row := range userEntityList {
			fmt.Println(row.Name, row.Age)
		}
	} else {
		log.Printf("user index error: %d", userIndexReponse.Err)
	}



	// UserView 请求
	userViewResponse, err := userClient.UserView(ctx, &user.UserViewRequest{Uid: 1})
	if err != nil {
		log.Printf("user view could not greet: %v", err)
	}

	if 0 == userViewResponse.Err {
		log.Printf("user view success: %s", userViewResponse.Msg)
		userEntity := userViewResponse.Data
		fmt.Println(userEntity.Name, userEntity.Age)
	} else {
		log.Printf("user view error: %d", userViewResponse.Err)
	}

	// UserPost 请求
	userPostReponse, err := userClient.UserPost(ctx, &user.UserPostRequest{Name: "big_cat", Password: "123456", Age: 29})
	if err != nil {
		log.Printf("user post could not greet: %v", err)
	}

	if 0 == userPostReponse.Err {
		log.Printf("user post success: %s", userPostReponse.Msg)
	} else {
		log.Printf("user post error: %d", userPostReponse.Err)
	}

	// UserDelete 请求
	userDeleteReponse, err := userClient.UserDelete(ctx, &user.UserDeleteRequest{Uid: 1})
	if err != nil {
		log.Printf("user delete could not greet: %v", err)
	}

	if 0 == userDeleteReponse.Err {
		log.Printf("user delete success: %s", userDeleteReponse.Msg)
	} else {
		log.Printf("user delete error: %d", userDeleteReponse.Err)
	}
}


启动服务/请求服务
go run server.go
# 新建一个窗口
go run client.go

```