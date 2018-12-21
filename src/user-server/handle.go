package user_server

import (
	"../github.com/micro/go-micro"
	"context"
	pb "./proto/user"
	"golang.org/x/crypto/go get -u golang.org/x/crypto/..."
)
const topic = "user.created"

type handler struct {
	Repo Repository
	TokenService Authable
	Publisher micro.Publisher
}

func (h *handler) Create(ctx context.Context,req *pb.User,resp *pb.Response)error{
	//hash password

}
