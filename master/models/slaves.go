package models

import (
	"fmt"

	"github.com/asdine/storm/q"
	pb "github.com/vladfr/arko/master/register"
)

type Slave struct {
	Id      int `storm:"id,increment"`
	Config  *pb.SlaveConfig
	Status  int `storm: "index"`
	Methods []string
}

func (db *DB) AddSlave(c *pb.SlaveConfig) *Slave {
	_, err := db.GetSlaveByConfig(c)

	if err != nil {
		// did not find a record, error is
		fmt.Println(fmt.Sprintf("Couldn't find record %s", err))
		s := &Slave{
			Config: c,
			Status: 1,
		}
		errS := db.Save(s)
		if errS != nil {
			fmt.Errorf("Cannot save slave to db")
		}
		return s
	}

	fmt.Println("Slave already registered, skipping.")
	return nil
}

func (db *DB) SaveSlave(s *Slave) error {
	return db.Save(s)
}

func (db *DB) GetSlaveByConfig(c *pb.SlaveConfig) (Slave, error) {
	var s Slave
	err := db.Select(q.Eq("Config", c)).First(&s)
	return s, err
}

func (db *DB) GetActiveSlaves() []Slave {
	var slaves []Slave
	err := db.Find("Status", 1, &slaves)
	if err != nil {
		fmt.Errorf("Cannot fetch active slaves")
		return nil
	}
	return slaves
}

func (db *DB) GetAllSlaves() []Slave {
	var slaves []Slave
	err := db.All(&slaves)
	if err != nil {
		fmt.Errorf("Cannot fetch all slaves")
		return nil
	}
	return slaves
}
