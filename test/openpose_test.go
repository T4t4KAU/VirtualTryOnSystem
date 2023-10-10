package test

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"govton/containers/openpose"
	"io/ioutil"
	"os"
	"testing"
)

func TestDockerOpenPoseExecCPU(t *testing.T) {
	currentDir, _ := os.Getwd()
	inputPath := currentDir + "/openpose_input"
	outputPath := currentDir + "/openpose_output"
	keypointPath := currentDir + "/openpose_keypoint"
	if _, _, err := openpose.ExecCPU(inputPath, outputPath, keypointPath); err != nil {
		t.Fatal(err)
	}
}

func DockerGPUOpenPose() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	containerConfig := container.Config{
		Image:      "openpose",
		Cmd:        []string{"bash", "run.sh"},
		Tty:        true,
		WorkingDir: "/openpose",
	}

	currentDir, _ := os.Getwd()
	hostConfig := container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: currentDir + "/openpose_input",
				Target: "/input",
			},
			{
				Type:   mount.TypeBind,
				Source: currentDir + "/openpose_output",
				Target: "/output",
			},
			{
				Type:   mount.TypeBind,
				Source: currentDir + "/openpose_keypoint",
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

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err = <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID,
		types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	return nil
}

func TestOpenPoseExec(t *testing.T) {
	currentDir, _ := os.Getwd()
	inputPath := currentDir + "/openpose_input"
	outputPath := currentDir + "/openpose_output"
	keypointPath := currentDir + "/openpose_keypoint"
	if _, _, err := openpose.Exec(inputPath, outputPath, keypointPath); err != nil {
		t.Fatal(err)
	}
}

func TestOpenPoseGenerate(t *testing.T) {
	image, _ := ioutil.ReadFile("./openpose_input/test.jpg")
	pose, _, err := openpose.Generate(image)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("images/pose_test.png", pose, os.ModePerm)
}
