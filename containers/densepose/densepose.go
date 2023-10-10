package densepose

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

type DensePoseContainer struct {
	id      string // 容器ID
	image   []byte // 图片
	address string // 地址
}

// 创建容器对象
func New() *DensePoseContainer {
	return &DensePoseContainer{}
}

// 设置容器IP地址
func (c *DensePoseContainer) set(id, address string) {
	c.id = id
	c.address = address
	containers.DensePoseAddress = address
}

// 生成图片
func (c *DensePoseContainer) Generate() ([]byte, error) {
	parse, _, err := process.DockerProcessHumanParse(c.image, c.address)
	return parse, err
}

// 启动预处理容器
func (c *DensePoseContainer) Run() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	containerConfig := container.Config{
		Image:      "venuns/densepose",
		Cmd:        []string{"python3", "server.py"},
		WorkingDir: "/densepose",
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
	log.Println("densepose conatiner start at", c.address)
	_, _ = cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	time.Sleep(3 * time.Second)

	return nil
}

func (c *DensePoseContainer) Store(image []byte) {
	c.image = image
}

func (c *DensePoseContainer) Address() string {
	return c.address
}

func (c *DensePoseContainer) Stop() error {
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
