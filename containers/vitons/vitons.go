package vitons

import (
	"context"
	"errors"
	"govton/containers"
	"govton/containers/agnostic"
	"govton/containers/densepose"
	"govton/containers/mask"
	"govton/containers/openpose"
	"govton/containers/parsing"
	"govton/containers/resize"
	"govton/process"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type VitonsContainer struct {
	// vitons推理必要的参数
	id                string
	address           string
	cloth             []byte
	image             []byte
	mask              []byte
	pose              []byte
	keypoint          string
	parse             []byte
	agnostic          []byte
	densepose         []byte
	PoseContainer     *densepose.DensePoseContainer
	MaskContainer     *mask.MaskContainer
	ParseContainer    *parsing.ParseContainer
	AgnosticContainer *agnostic.AgnosticContainer
}

func New() *VitonsContainer {
	c := &VitonsContainer{}
	c.PoseContainer = densepose.New()
	c.MaskContainer = mask.New()
	c.ParseContainer = parsing.New()

	return c
}

// 设置容器ID和地址
func (c *VitonsContainer) set(id, address string) {
	c.id = id
	c.address = address
	containers.VitonsAddress = address
}

func (c *VitonsContainer) Run() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	containerConfig := container.Config{
		Image:      "venuns/vitons",
		Cmd:        []string{"python3", "server.py"},
		WorkingDir: "/HR_VITON",
	}

	hostConfig := container.HostConfig{
		AutoRemove: true,
		ShmSize:    int64(8) * 1024 * 1024 * 1024,
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
	log.Println("vitons container start at", c.address)
	_, _ = cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	time.Sleep(5 * time.Second)

	return nil
}

func (c *VitonsContainer) Generate() ([]byte, error) {
	return process.DockerVitonsGenerate(c.cloth, c.mask, c.image, c.parse,
		c.pose, c.keypoint, c.agnostic, c.densepose, c.address)
}

func (c *VitonsContainer) Read(image, cloth []byte) {
	c.image = image
	c.cloth = cloth
}

func (c *VitonsContainer) Store(cloth []byte, mask []byte, image []byte, parse []byte,
	pose []byte, keypoint string, agnostic []byte, densepose []byte) {
	c.cloth = cloth
	c.mask = mask
	c.image = image
	c.parse = parse
	c.pose = pose
	c.keypoint = keypoint
	c.agnostic = agnostic
	c.densepose = densepose
}

func (c *VitonsContainer) Address() string {
	return c.address
}

func (c *VitonsContainer) Stop() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.41"))
	if err != nil {
		return err
	}

	timeout := 500
	options := container.StopOptions{
		Timeout: &timeout,
	}
	err = cli.ContainerStop(context.Background(), c.id, options)
	if err != nil {
		return err
	}
	if err = c.MaskContainer.Stop(); err != nil {
		return err
	}
	if err = c.ParseContainer.Stop(); err != nil {
		return err
	}
	if err = c.PoseContainer.Stop(); err != nil {
		return err
	}
	if err = c.AgnosticContainer.Stop(); err != nil {
		return err
	}
	return nil
}
func (c *VitonsContainer) Process() error {
	if c.PoseContainer == nil || c.MaskContainer == nil ||
		c.ParseContainer == nil {
		return errors.New("invalid container")
	}
	if err := c.PoseContainer.Run(); err != nil {
		return err
	}
	if err := c.MaskContainer.Run(); err != nil {
		return err
	}
	if err := c.ParseContainer.Run(); err != nil {
		return err
	}
	return nil
}

func (c *VitonsContainer) AutoGenerate() ([]byte, error) {
	if c.PoseContainer == nil || c.MaskContainer == nil ||
		c.ParseContainer == nil || c.AgnosticContainer == nil {
		return nil, errors.New("invalid container")
	}

	poseImage, keypoint, err := openpose.Generate(c.image)
	if err != nil {
		return nil, err
	}
	c.pose, c.keypoint = poseImage, keypoint

	c.PoseContainer.Store(c.image)
	denseposeImage, err := c.PoseContainer.Generate()
	if err != nil {
		return nil, err
	}
	c.densepose = denseposeImage

	c.MaskContainer.Store(c.cloth)
	maskImage, err := c.MaskContainer.Generate()
	if err != nil {
		return nil, err
	}
	c.mask = maskImage

	c.ParseContainer.Store(c.image)
	parseImage, err := c.ParseContainer.Generate()
	if err != nil {
		return nil, err
	}
	c.parse = parseImage

	c.AgnosticContainer.Store(c.image, c.cloth, c.mask, c.parse, c.densepose, c.pose, keypoint)
	c.agnostic, err = c.AgnosticContainer.Generate()
	if err != nil {
		return nil, err
	}
	result, err := c.Generate()
	if err != nil {
		return nil, err
	}
	return result, err
}

func Generate(image, cloth []byte) ([]byte, error) {
	image, cloth, err := resize.Generate(image, cloth)
	if err != nil {
		return nil, err
	}
	pose, keypoint, err := openpose.Generate(image)
	if err != nil {
		return nil, err
	}

	mask, err := process.DockerProcessClothSegmentation(
		cloth, containers.ClothMaskAddress)
	if err != nil {
		return nil, err
	}

	parse, _, err := process.DockerProcessHumanParse(
		image, containers.HumanParseAddress)
	if err != nil {
		return nil, err
	}

	densepose, err := process.DockerProcessDensePose(
		image, containers.DensePoseAddress)
	if err != nil {
		return nil, err
	}

	agnostic, err := agnostic.Generate(image, cloth, mask,
		parse, densepose, pose, string(keypoint))
	if err != nil {
		return nil, err
	}

	result, err := process.DockerVitonsGenerate(cloth, mask, image, parse,
		pose, keypoint, agnostic, densepose, containers.VitonsAddress)
	if err != nil {
		return nil, err
	}
	return result, nil
}
