package rdb

import (
	"log"
	"os"
	"testing"
	"time"

	gorethink "gopkg.in/dancannon/gorethink.v2"

	"github.com/deskr/gopkg/dock"
	docker "github.com/fsouza/go-dockerclient"
)

func TestMain(m *testing.M) {
	conf := dock.Config{}
	conf.Image = "rethinkdb:latest"
	conf.Tty = true
	conf.OpenStdin = true

	// Using "Docker for Mac" - so no connection from host -> container :(
	conf.ExposedPorts = map[docker.Port]struct{}{
		"28015/tcp": {},
	}
	conf.PortBindings = map[docker.Port][]docker.PortBinding{
		"28015/tcp": []docker.PortBinding{
			{
				HostPort: "29017",
			},
		},
	}
	_, closer, err := dock.Run(conf)
	if err != nil {
		log.Fatalf("Failed to run RethinkDB container: %+v", err)
		return
	}

	res := m.Run()

	closer()

	os.Exit(res)
}

func TestOpenSession(t *testing.T) {
	_, err := OpenSession(gorethink.ConnectOpts{
		Address:  "localhost:29017",
		Database: "test",
	}, time.Second*5)
	if err != nil {
		t.Errorf("Failed to open session: %v", err)
		return
	}
}

// TestOpenSessionNoWait works only if it get's tested after another connection has been made
func TestOpenSessionNoWait(t *testing.T) {
	_, err := OpenSession(gorethink.ConnectOpts{
		Address:  "localhost:29017",
		Database: "test",
	}, 0)
	if err != nil {
		t.Errorf("Failed to open session: %v", err)
		return
	}
}
