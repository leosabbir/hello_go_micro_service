package main

import (
	"context"
	"log"
	"micros/consignment-service/proto/consignment"
	"micros/vessel-service/proto/vessel"

	mgo "gopkg.in/mgo.v2"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	session      *mgo.Session
	vesselClient vessel.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *consignment.Consignment, res *consignment.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vessel.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	req.VesselId = vesselResponse.Vessel.Id
	// Save our consignment
	err = repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) ReadConsignments(ctx context.Context, req *consignment.GetRequest, res *consignment.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	consignments, err := repo.GetConsignments()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
