package main

import (
	"micros/consignment-service/proto/consignment"

	"gopkg.in/mgo.v2"
)

const (
	dbName                = "micros"
	consignmentCollection = "consignments"
)

// Repository interface for repository
type Repository interface {
	Create(*consignment.Consignment) error
	GetConsignments() ([]*consignment.Consignment, error)
	Close()
}

// ConsignmentRepository - Dummy repository, this simulates the use of a datastore of
// some kind.
type ConsignmentRepository struct {
	session *mgo.Session
}

// Create creates a new consignment and stores in the repository
func (repo *ConsignmentRepository) Create(consignment *consignment.Consignment) error {
	return repo.collection().Insert(consignment)
}

// GetConsignments returns all the created consignments
func (repo *ConsignmentRepository) GetConsignments() ([]*consignment.Consignment, error) {
	var consignments []*consignment.Consignment
	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}

// Close closes the database session after each query has run
func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(consignmentCollection)
}
