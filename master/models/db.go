package models

import (
	"fmt"

	"github.com/asdine/storm"
	"github.com/asdine/storm/v3/codec/protobuf"
	pb "github.com/vladfr/arko/master/register"
)

type Datastore interface {
	GetActiveSlaves() []Slave
	GetAllSlaves() []Slave
	AddSlave(*pb.SlaveConfig) *Slave
	GetSlaveByConfig(*pb.SlaveConfig) (Slave, error)
	SaveSlave(*Slave) error
}

type DB struct {
	*storm.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := storm.Open(dataSourceName, storm.Codec(protobuf.Codec))
	if err != nil {
		fmt.Printf("Cannot open database at %s", dataSourceName)
		panic(err)
	}
	return &DB{db}, nil
}
