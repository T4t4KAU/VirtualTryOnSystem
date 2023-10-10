package resize

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Generate(image, cloth []byte) ([]byte, []byte, error) {
	inputPath, _ := filepath.Abs("../test/resize_input")
	outputPath, _ := filepath.Abs("../test/resize_output")
	ioutil.WriteFile(inputPath+"/image/test.jpg", image, os.ModePerm)
	ioutil.WriteFile(inputPath+"/cloth/test.jpg", cloth, os.ModePerm)
	return Exec(inputPath, outputPath)
}

func Exec(inputPath, outputPath string) ([]byte, []byte, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, nil, err
	}
	defer cli.Close()

	containerConfig := container.Config{
		Image:      "venuns/resize:viton",
		Cmd:        []string{"python", "resize.py"},
		WorkingDir: "/resize",
	}

	hostConfig := container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: inputPath,
				Target: "/resize/input",
			},
			{
				Type:   mount.TypeBind,
				Source: outputPath,
				Target: "/resize/output",
			},
		},
		AutoRemove: true,
	}

	log.Println("resize processing")
	resp, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig,
		nil, nil, "")
	if err != nil {
		return nil, nil, err
	}
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return nil, nil, err
	}
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err = <-errCh:
		if err != nil {
			return nil, nil, err
		}
	case <-statusCh:
	}

	resizeImage, _ := ioutil.ReadFile(outputPath + "/image/test.jpg")
	resizeCloth, _ := ioutil.ReadFile(outputPath + "/cloth/test.jpg")
	return resizeImage, resizeCloth, nil
}
