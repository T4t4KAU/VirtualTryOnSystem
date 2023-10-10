package agnostic

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

type AgnosticContainer struct {
	id        string
	address   string
	cloth     []byte
	mask      []byte
	image     []byte
	densepose []byte
	parse     []byte
	pose      []byte
	keypoint  string
}

func New() *AgnosticContainer {
	return &AgnosticContainer{}
}

func (c *AgnosticContainer) set(id, address string) {
	c.id = id
	c.address = address
	containers.ParseAgnosticAddress = address
}

func (c *AgnosticContainer) Run() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	containerConfig := container.Config{
		Image:      "venuns/parse-agnostic",
		Cmd:        []string{"python3", "server.py"},
		WorkingDir: "/parse_agnostic",
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
	log.Println("parse-agnostic conatiner start at", c.address)
	_, _ = cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	time.Sleep(5 * time.Second)

	return nil
}

func (c *AgnosticContainer) Generate() ([]byte, error) {
	return process.DockerProcessParseAgnostic(c.cloth, c.mask,
		c.image, c.parse, c.pose, c.keypoint, c.densepose, c.address)
}

func (c *AgnosticContainer) Store(image, cloth, mask, parse, densepose, pose []byte, keypoint string) {
	c.image = image
	c.cloth = cloth
	c.mask = mask
	c.pose = pose
	c.keypoint = keypoint
	c.densepose = densepose
	c.parse = parse
}

func (c *AgnosticContainer) Address() string {
	return c.address
}

func (c *AgnosticContainer) Stop() error {
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
	log.Println("densepose container stop at", c.address)
	return nil
}
