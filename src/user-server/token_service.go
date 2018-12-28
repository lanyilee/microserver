package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	pb "user-server/src/proto/user"
)

var privateKey = []byte("`xs#a_1-!")

type Authable interface {
	Decode(tokenStr string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type TokenService struct {
	repo Repository
}

func (srv *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (srv *TokenService) Encode(user *pb.User) (string, error) {
	// one days
	expireTime := time.Now().Add(time.Hour * 24).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			Issuer:    "go.micro.srv.user",
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return jwtToken.SignedString(privateKey)
}
