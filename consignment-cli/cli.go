package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"micros/consignment-service/proto/consignment"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

const (
	//address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*consignment.Consignment, error) {
	var c consignment.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &c)
	return &c, err
}

func main() {
	cmd.Init()
	client := consignment.NewShippingServiceClient("go.micro.service.consignment", microclient.DefaultClient)

	// Contact the server and print out its Response
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	c, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	ctx := context.Background()
	r, err := client.CreateConsignment(ctx, c)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Create: %t", r.Created)

	r, err = client.ReadConsignments(ctx, &consignment.GetRequest{})
	if err != nil {
		log.Fatalf("Could not read consignments: %v", err)
	}
	for _, v := range r.GetConsignments() {
		log.Println(v)
	}
}
