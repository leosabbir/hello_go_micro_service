// user-service/handler.go
package main

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"log"
	pb "micros/user-service/proto/user"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in with: ", req.Email, req.Passwords)
	user, err := srv.repo.GetByEmailAndPassword(req)
	log.Println("User Read: %v", user)
	if err != nil {
		return err
	}

	// Compare given password against hashed password in database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Passwords), []byte(req.Passwords)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}

	res.Token = token
	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	// Generates a hased password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Passwords), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Passwords = string(hashedPass)
	if err := srv.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
