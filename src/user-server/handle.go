package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	_ "github.com/micro/go-plugins/broker/nats"
	"golang.org/x/crypto/bcrypt"
	"log"
	pb "user-server/src/proto/user"
)

const topic = "user.created"

type handler struct {
	Repo         Repository
	TokenService Authable
	Publisher    micro.Publisher
}

func (h *handler) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	//hash password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
		return err
	}
	req.Password = string(hashedPwd)
	err = h.Repo.Create(req)
	if err != nil {
		log.Panic(err)
		return err
	}
	resp.User = req
	// publish user
	if err := h.Publisher.Publish(ctx, req); err != nil {
		return err
	}
	return nil
}

func (h *handler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	u, err := h.Repo.GetByEmail(req.Email)
	if err != nil {
		log.Panic(err)
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		log.Panic(err)
		return err
	}
	//to generate token
	t, err := h.TokenService.Encode(u)
	resp.Token = t
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	claim, err := h.TokenService.Decode(req.Token)
	if err != nil {
		log.Panic(err)
		return err
	}
	if claim.User.Id == "" {
		return errors.New("1", "invalid user", 101)
	}
	resp.Valid = true
	return nil
}

func (h *handler) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	user, err := h.Repo.Get(req.Id)
	if err != nil {
		log.Panic(err)
		return err
	}
	resp.User = user
	return nil
}

func (h *handler) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	users, err := h.Repo.GetAll()
	if err != nil {
		log.Panic(err)
		return err
	}
	resp.Users = users
	return nil
}
