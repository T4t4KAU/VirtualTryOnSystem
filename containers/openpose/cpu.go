package openpose

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

func Generate(image []byte) ([]byte, string, error) {
	stamp := fmt.Sprintf("%d", time.Now().Unix())
	root, _ := filepath.Abs("../temp/openpose")
	inputPath := root + "/" + stamp + "_input"
	outputPath := root + "/" + stamp + "_output"
	keypointPath := root + "/" + stamp + "_keypoint"

	os.Mkdir(inputPath, os.ModePerm)
	os.Mkdir(outputPath, os.ModePerm)
	os.Mkdir(keypointPath, os.ModePerm)

	err := ioutil.WriteFile(inputPath+"/test.jpg", image, os.ModePerm)
	if err != nil {
		return nil, "", err
	}
	pose, keypoint, err := ExecCPU(inputPath, outputPath, keypointPath)

	os.RemoveAll(inputPath)
	os.RemoveAll(outputPath)
	os.RemoveAll(keypointPath)

	return pose, keypoint, err
}

func ExecCPU(inputPath string, outputPath string, keypointPath string) ([]byte, string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, "", err
	}
	defer cli.Close()

	containerConfig := container.Config{
		Image: "nephilimboy/openpose-cpu",
		Cmd: []string{"--image_dir", "/input",
			"--write_images", "/output",
			"--write_json", "/keypoint",
			"--display", "0",
			"--hand", "--disable_blending",
		},
	}

	hostConfig := container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: inputPath,
				Target: "/input",
			},
			{
				Type:   mount.TypeBind,
				Source: outputPath,
				Target: "/output",
			},
			{
				Type:   mount.TypeBind,
				Source: keypointPath,
				Target: "/keypoint",
			},
		},
		AutoRemove: true,
	}

	log.Println("openpose-cpu processing")
	resp, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig,
		nil, nil, "")
	if err != nil {
		return nil, "", err
	}
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return nil, "", err
	}
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err = <-errCh:
		if err != nil {
			return nil, "", err
		}
	case <-statusCh:
	}

	image, _ := ioutil.ReadFile(outputPath + "/test_rendered.png")
	keypoint, _ := ioutil.ReadFile(keypointPath + "/test_keypoints.json")

	return image, string(keypoint), err
}
