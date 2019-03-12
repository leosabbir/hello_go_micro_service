package main

import (
	"log"
	"micros/common"
	"micros/consignment-service/proto/consignment"
	"micros/constants"
	"os"

	vessel "micros/vessel-service/proto/vessel"

	micro "github.com/micro/go-micro"
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = constants.DefaultHost
	}

	session, err := common.CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.service.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vessel.NewVesselServiceClient("go.micro.service.vessel", srv.Client())

	srv.Init()

	consignment.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
