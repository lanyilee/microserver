package main

import (
	pb "../user-server/src/proto/user"
	microClinet "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
	"log"
	"os"
)

func main() {
	cmd.Init()
	//create user-service micro server client
	client := pb.NewUserServiceClient("go.micro.srv.user", microClinet.DefaultClient)
	//test user
	user := &pb.User{
		Name:     "lanyi",
		Email:    "lanyilee@qq.com",
		Password: "admin",
		Company:  "google",
	}
	resp, err := client.Create(context.TODO(), user)
	if err != nil {
		log.Panic(err)
	}
	log.Println("created:", resp.User.Id)
	allResp, err := client.GetAll(context.Background(), &pb.Request{})
	for i, user := range allResp.Users {
		log.Println("user_%d:%v\n", i, user)
	}
	authResp, err := client.Auth(context.TODO(), user)
	log.Println("token:", authResp.Token)
	os.Exit(0)
}
