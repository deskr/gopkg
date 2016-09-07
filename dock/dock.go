package dock

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/fsouza/go-dockerclient"
)

var (
	Address = `unix:///var/run/docker.sock`
)

type Config struct {
	docker.Config
	docker.HostConfig
}

// Run a docker container
func Run(config Config) (ip string, closer func(), err error) {
	client, err := docker.NewClient(Address)
	if err != nil {
		return
	}

	c, err := client.CreateContainer(docker.CreateContainerOptions{
		Config:     &config.Config,
		HostConfig: &config.HostConfig,
	})
	if err != nil {
		return
	}

	closer = func() {
		if err := client.RemoveContainer(docker.RemoveContainerOptions{
			ID:    c.ID,
			Force: true,
		}); err != nil {
			log.Printf("cannot remove container: %v", err)
		}
	}

	if err = client.StartContainer(c.ID, &config.HostConfig); err != nil {
		closer()
		return
	}

	// wait for container to wake up
	if err = waitStarted(client, c.ID, 5*time.Second); err != nil {
		closer()
		return
	}
	if c, err = client.InspectContainer(c.ID); err != nil {
		closer()
		return
	}

	ip = strings.TrimSpace(c.NetworkSettings.IPAddress)
	return
}

// waitStarted waits for container to start for the maxWait time.
func waitStarted(client *docker.Client, id string, maxWait time.Duration) error {
	done := time.Now().Add(maxWait)
	for time.Now().Before(done) {
		c, err := client.InspectContainer(id)
		if err != nil {
			break
		}
		if c.State.Running {
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("Failed to start container %s for %v", id, maxWait)
}

// WaitReachable waits for a successful connection
func WaitReachable(host string, port int, maxWait time.Duration) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	done := time.Now().Add(maxWait)
	for time.Now().Before(done) {
		c, err := net.DialTimeout("tcp", addr, time.Second*1)
		if err == nil {
			c.Close()
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("Failed to connect %v for %v", addr, maxWait)
}
