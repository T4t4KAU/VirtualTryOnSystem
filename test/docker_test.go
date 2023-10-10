package test

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"govton/containers"
	"os"
	"testing"
)

func TestDockerAPI(t *testing.T) {
	if err := ListDockerImages(); err != nil {
		t.Fatal(err)
	}
}

func ListDockerImages() error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return err
	}
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID, container.Image)
	}
	return nil
}

func TestDockerMount(t *testing.T) {
	if err := DockerMountFloder(); err != nil {
		t.Fatal(err)
	}
}

func DockerMountFloder() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()
	currentDir, _ := os.Getwd()

	containerConfig := container.Config{
		Image: "alpine",
		Cmd:   []string{"ls", "/"},
	}
	hostConfig := container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: currentDir + "/testdir",
				Target: "/test",
			},
		},
	}
	resp, err := cli.ContainerCreate(ctx, &containerConfig,
		&hostConfig, nil, nil, "")
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

func CheckDockerRunning(id string) (error, bool) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err, false
	}
	defer cli.Close()
	insepct, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		return err, false
	}
	return nil, insepct.State.Running
}

func TestContainerIPAddress(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.41"))
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	address, err := containers.ContainerIPAddress(cli, "3bfc791c44ca")
	if err != nil {
		t.Fatal(err)
	}
	println(address)
}
