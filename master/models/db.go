package models

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/v3/codec/protobuf"
	pb "github.com/vladfr/arko/master/register"
)

type Datastore interface {
	ActiveSlaves() []Slave
	AddSlave(*pb.SlaveConfig)
	GetSlaveByConfig(*pb.SlaveConfig) (Slave, error)
	SaveSlave(*Slave) error
}

type DB struct {
	*storm.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := storm.Open(dataSourceName, storm.Codec(protobuf.Codec))
	if err != nil {
		panic("Cannot open database file")
	}
	return &DB{db}, nil
}
