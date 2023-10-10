package openpose

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"io/ioutil"
)

func Exec(inputPath, outputPath, keypointPath string) ([]byte, string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, "", err
	}
	defer cli.Close()

	containerConfig := container.Config{
		Image:      "openpose",
		Cmd:        []string{"bash", "run.sh"},
		Tty:        true,
		WorkingDir: "/openpose",
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
		Resources: container.Resources{
			DeviceRequests: []container.DeviceRequest{
				{
					Driver:       "nvidia",
					Count:        0,
					DeviceIDs:    []string{"0"},
					Capabilities: [][]string{{"gpu"}},
					Options:      nil,
				},
			},
		},
		AutoRemove: true,
	}

	resp, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig,
		nil, nil, "")
	if err != nil {
		return nil, "", err
	}
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return nil, "", err
	}
	cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	poseImage, _ := ioutil.ReadFile(outputPath + "/test_rendered.png")
	poseJson, _ := ioutil.ReadFile(keypointPath + "/test_keypoints.json")
	return poseImage, string(poseJson), nil
}
