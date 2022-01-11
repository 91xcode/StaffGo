package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	pb "code.be.staff.com/staff/StaffGo/grpc/go_grpc_mysql/user"
	"fmt"
)

func main() {

	r := gin.Default()

	r.GET("/rpc/get-user-list", func(c *gin.Context) {
		sayHello(c)
	})

	// Run http server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func sayHello(c *gin.Context) {

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8084", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	name := c.DefaultQuery("name","")
	mobile := c.DefaultQuery("mobile","")
	req := &pb.RequestUser{Name: name,Mobile:mobile}
	res, err := client.UserList(c, req)

	fmt.Println("res: %v", res.User)
	

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)

}




