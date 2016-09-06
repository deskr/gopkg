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

func Run(config Config) (ip string, closer func(), err error) {
	client, err := docker.NewClient(Address)
	if err != nil {
		return "", nil, err
	}

	c, err := client.CreateContainer(docker.CreateContainerOptions{
		Config:     &config.Config,
		HostConfig: &config.HostConfig,
	})
	if err != nil {
		return "", nil, err
	}

	closer = func() {
		if err := client.RemoveContainer(docker.RemoveContainerOptions{
			ID:    c.ID,
			Force: true,
		}); err != nil {
			log.Printf("cannot remove container: %v", err)
		}
	}

	err = client.StartContainer(c.ID, &config.HostConfig)
	if err != nil {
		closer()
		return "", nil, err
	}

	// wait for container to wake up
	if err := waitStarted(client, c.ID, 5*time.Second); err != nil {
		closer()
		return "", nil, err
	}
	if c, err = client.InspectContainer(c.ID); err != nil {
		closer()
		return "", nil, err
	}

	return strings.TrimSpace(c.NetworkSettings.IPAddress), closer, nil
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

func WaitReachable(addr string, maxWait time.Duration) error {
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
