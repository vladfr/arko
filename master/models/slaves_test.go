package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	pb "github.com/vladfr/arko/master/register"
)

func TestMain(m *testing.M) {
	code := m.Run()
	teardown()
	os.Exit(code)
}

func getDB() (*DB, func()) {
	dir, _ := ioutil.TempDir(os.TempDir(), "arkotest")

	tdb, err := NewDB(filepath.Join(dir, "arkotest.db"))
	if err != nil {
		fmt.Println("Cannot open temp db")
		panic(err)
	}
	return tdb, func() {
		tdb.Close()
		os.RemoveAll(dir)
	}
}

func teardown() {
}

func TestAddSlave(t *testing.T) {
	tdb, close := getDB()
	defer close()

	slavec := &pb.SlaveConfig{
		Host:  "testhost",
		Port:  80,
		Token: "bla",
	}

	s := tdb.AddSlave(slavec)
	if s == nil {
		t.Error("Slave not saved")
	}

	if s.Config.Host != "testhost" {
		t.Errorf("Something is wrong, slave host is %s, should be testhost", s.Config.Host)
	}
	if s.Status != 1 {
		t.Errorf("Slave saved but status %d != 1", s.Status)
	}
}

func TestGetSlaveByConfig(t *testing.T) {
	tdb, close := getDB()
	defer close()

	slavec := &pb.SlaveConfig{
		Host:  "testhost",
		Port:  80,
		Token: "bla",
	}

	tdb.AddSlave(slavec)
	s, err := tdb.GetSlaveByConfig(slavec)
	if err != nil {
		t.Error(err)
	}

	if s.Config.Host != "testhost" {
		t.Errorf("Something is wrong, slave host is %s, should be testhost", s.Config.Host)
	}
}

func TestAddSlaveTwice(t *testing.T) {
	tdb, close := getDB()
	defer close()

	slavec := &pb.SlaveConfig{
		Host:  "testhost",
		Port:  80,
		Token: "bla",
	}

	tdb.AddSlave(slavec)
	tdb.AddSlave(slavec)

	slaves := tdb.GetAllSlaves()
	if len(slaves) > 1 {
		t.Error("Added the same slave twice, but found more than one slave")
	}

	for _, s := range tdb.GetAllSlaves() {
		if s.Config.Host != "testhost" {
			t.Errorf("Something is wrong, slave host is %s, should be testhost", s.Config.Host)
		}
		if s.Status != 1 {
			t.Errorf("Slave saved but status %d != 1", s.Status)
		}
	}
}
