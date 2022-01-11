package main

import (
	"code.be.staff.com/staff/StaffGo/grpc/go_grpc_mysql/common"
	"code.be.staff.com/staff/StaffGo/grpc/go_grpc_mysql/models"
	user_pb "code.be.staff.com/staff/StaffGo/grpc/go_grpc_mysql/user"
	"context"
	"fmt"
	"github.com/go-xorm/xorm"
	"google.golang.org/grpc"
	"log"
	"net"
)



//userservice的实现
type UserServiceImpl struct {
	//因为我们要使用到xorm操作数据库 所以要定义哦 在注册的时候我们会将数据库引擎注册进去
	Engine *xorm.Engine
}
//userservice实现的方法  UserList()里面具体的参数可以直接参考User.pb.go文件里面的UserList()方法
func (us *UserServiceImpl) UserList(ctx context.Context,user *user_pb.RequestUser) (*user_pb.ResponseUser,error) {
	//返回的是 user2.ResponseUser类型
	return &user_pb.ResponseUser{
		User:us.GetUser(user.Name),
	},nil
}
//userservice实现的方法
func (us *UserServiceImpl) GetUser(name string) []*user_pb.User {

	//记得去创建user的模型
	var userList []models.User
	err := us.Engine.Where("id > ?", 0).Find(&userList)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(userList)
	fmt.Println("===========================================")

	//因为返回的是[]*user2.User
	userList1 := make([]*user_pb.User,0)
	//遍历数据库查询结果 然后重新塞入到新的切片当中去
	for _,u := range userList {
		userList1 = append(userList1,&user_pb.User{
			Name:u.Name,
			Age:u.Age,
			Mobile:u.Mobile,
		})
	}

	return userList1

	//users := make([]*user2.User,0)
	//for i:=1;i<=10;i++ {
	//    n := strconv.Itoa(i)
	//    users = append(users,&user2.User{
	//        Name:n,
	//        Age:int64(i),
	//    })
	//}
	//return users;

}

func main(){
	//服务端起来一个tcp服务 监听端口8084
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		fmt.Println("server error 500",err.Error())
		return
	}
	//实现一个grpc
	server := grpc.NewServer()
	//数据库引擎
	engine :=  common.NewMysqlEngine()
	//注册service的时候我们就将数据库引擎传递进去 这样在service里面就可以操作数据库啦
	user_pb.RegisterUserServiceServer(server,&UserServiceImpl{Engine:engine})
	//启动服务   这个地方我们不需要for死循环 grpc当中会自动帮助我们实时监听
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}