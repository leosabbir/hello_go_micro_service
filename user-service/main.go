package main

import (
	"fmt"
	micro "github.com/micro/go-micro"
	"log"
	pb "micros/user-service/proto/user"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc. This will check for changes and migrate them each time this service is restarted.
	db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}
	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
