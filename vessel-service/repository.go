package main

import (
	"errors"
	pb "micros/vessel-service/proto/vessel"

	mgo "gopkg.in/mgo.v2"
)

const (
	dbName           = "micros"
	vesselCollection = "vessels"
)

// Repository stores the vessels
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

// VesselRepository actual repository
type VesselRepository struct {
	session *mgo.Session
}

// FindAvailable finds vessels that matches the Specification
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessels []*pb.Vessel
	err := repo.collection().Find(nil).All(&vessels)
	if err != nil {
		return nil, err
	}
	for _, vessel := range vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

// Close closes the database session after each query has run
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

// Create creates new vessel in the database
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
