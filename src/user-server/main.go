package main

import (
	"github.com/micro/go-micro"
	"log"
	pb "user-server/src/proto/user"
)

func main() {
	config, err := ReadConfig("src/config.conf")
	if err != nil {
		log.Panic(err)
	}
	db := config.CreateConnection()
	defer db.Close()
	repo := &UserRepository{db: db}
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)
	srv.Init()
	// pubSub := s.Server().Options().Broker
	publisher := micro.NewPublisher(topic, srv.Client())
	t := TokenService{repo: repo}

	pb.RegisterUserServiceHandler(srv.Server(), &handler{repo, &t, publisher})
	if err = srv.Run(); err != nil {
		log.Panic(err)
	}

}
