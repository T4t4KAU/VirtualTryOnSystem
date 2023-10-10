package resolution

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Generate(image []byte) ([]byte, error) {
	stamp := fmt.Sprintf("%d", time.Now().Unix())
	root, _ := filepath.Abs("../temp/resolution")
	inputPath := root + "/" + stamp + "_input"
	outputPath := root + "/" + stamp + "_output"

	os.Mkdir(inputPath, os.ModePerm)
	os.Mkdir(outputPath, os.ModePerm)

	ioutil.WriteFile(inputPath+"/test.jpg", image, os.ModePerm)
	ioutil.WriteFile(outputPath+"/test.jpg", image, os.ModePerm)

	res, err := ExecCPU(inputPath, outputPath)

	os.RemoveAll(inputPath)
	os.RemoveAll(outputPath)

	return res, err
}

func ExecCPU(inputPath string, outputPath string) ([]byte, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	containerConfig := container.Config{
		Image:      "venuns/super-resolution",
		Cmd:        []string{"bash", "test.sh"},
		WorkingDir: "/SRGAN",
	}

	hostConfig := container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: inputPath,
				Target: "/SRGAN/input",
			},
			{
				Type:   mount.TypeBind,
				Source: outputPath,
				Target: "/SRGAN/output",
			},
		},
		AutoRemove: true,
	}

	log.Println("super-resolution processing")
	resp, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig,
		nil, nil, "")
	if err != nil {
		return nil, err
	}
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return nil, err
	}
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err = <-errCh:
		if err != nil {
			return nil, err
		}
	case <-statusCh:
	}

	image, _ := ioutil.ReadFile(outputPath + "/test.jpg")

	return image, err
}
