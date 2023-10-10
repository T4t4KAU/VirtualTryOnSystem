package containers

import (
	"context"

	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

var (
	DensePoseAddress     string
	HumanParseAddress    string
	ClothMaskAddress     string
	VitonsAddress        string
	ParseAgnosticAddress string
)

type Config struct {
	image  string
	mounts map[string]string
}

func Mount(config Config) []mount.Mount {
	mounts := make([]mount.Mount, 0)
	for source, target := range config.mounts {
		mounts = append(mounts, mount.Mount{Type: mount.TypeBind, Source: source, Target: target})
	}
	return mounts
}

// get container address
func ContainerIPAddress(cli *client.Client, containerID string) (string, error) {
	inspect, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return "", err
	}
	address := inspect.NetworkSettings.IPAddress
	return address, nil
}
