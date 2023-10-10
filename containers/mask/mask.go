package mask

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"govton/containers"
	"govton/process"
	"log"
	"time"
)

type MaskContainer struct {
	id      string
	address string
	cloth   []byte
}

func New() *MaskContainer {
	return &MaskContainer{}
}

func (c *MaskContainer) set(id, address string) {
	c.id = id
	c.address = address
	containers.ClothMaskAddress = address
}

func (c *MaskContainer) Run() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	containerConfig := container.Config{
		Image:      "venuns/clothmask",
		Cmd:        []string{"python3", "server.py"},
		WorkingDir: "/cloth-segmentation",
	}

	hostConfig := container.HostConfig{
		AutoRemove: true,
	}

	resp, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig,
		nil, nil, "")
	if err != nil {
		return err
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	address, err := containers.ContainerIPAddress(cli, resp.ID)
	if err != nil {
		return err
	}

	c.set(resp.ID, address)
	log.Println("clothmask conatiner start at", c.address)
	_, _ = cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	time.Sleep(3 * time.Second)

	return nil
}

func (c *MaskContainer) Generate() ([]byte, error) {
	return process.DockerProcessClothSegmentation(c.cloth, c.address)
}

func (c *MaskContainer) Stop() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.41"))
	if err != nil {
		return err
	}

	timeout := 20
	options := container.StopOptions{
		Timeout: &timeout,
	}
	err = cli.ContainerStop(context.Background(), c.id, options)
	if err != nil {
		return err
	}
	log.Println("clothmask container stop at", c.address)
	return nil
}

func (c *MaskContainer) Store(cloth []byte) {
	c.cloth = cloth
}

func (c *MaskContainer) Address() string {
	return c.address
}
