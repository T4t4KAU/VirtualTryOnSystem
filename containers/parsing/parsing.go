package parsing

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

type ParseContainer struct {
	id      string
	address string
	image   []byte
}

func New() *ParseContainer {
	return &ParseContainer{}
}

func (c *ParseContainer) set(id, address string) {
	c.id = id
	c.address = address
	containers.HumanParseAddress = address
}

func (c *ParseContainer) Generate() ([]byte, error) {
	parse, _, err := process.DockerProcessHumanParse(c.image, c.address)
	return parse, err
}

func (c *ParseContainer) Run() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	containerConfig := container.Config{
		Image:      "venuns/humanparse",
		Cmd:        []string{"python3", "server.py"},
		WorkingDir: "/human_parse",
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
	log.Println("human-parse container start at", c.address)
	_, _ = cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	time.Sleep(30 * time.Second)

	return nil
}

func (c *ParseContainer) Address() string {
	return c.address
}

func (c *ParseContainer) Store(image []byte) {
	c.image = image
}

func (c *ParseContainer) Stop() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.41"))
	if err != nil {
		return err
	}

	timeout := 10
	options := container.StopOptions{
		Timeout: &timeout,
	}
	err = cli.ContainerStop(context.Background(), c.id, options)
	if err != nil {
		return err
	}
	log.Println("parsing conatiner stop at", c.address)

	return nil
}
