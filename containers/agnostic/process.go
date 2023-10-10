package agnostic

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

func Generate(image, cloth, mask, parse, densepose, pose []byte, keypoint string) ([]byte, error) {
	root, err := filepath.Abs("../temp/agnostic")
	if err != nil {
		return nil, err
	}
	stamp := fmt.Sprintf("%d", time.Now().Unix())
	inputPath := root + "/" + stamp + "_input"
	outputPath := root + "/" + stamp + "_output"

	os.Mkdir(inputPath, os.ModePerm)
	os.Mkdir(outputPath, os.ModePerm)

	os.Mkdir(inputPath+"/image", os.ModePerm)
	os.Mkdir(inputPath+"/cloth", os.ModePerm)
	os.Mkdir(inputPath+"/cloth-mask", os.ModePerm)
	os.Mkdir(inputPath+"/image-densepose", os.ModePerm)
	os.Mkdir(inputPath+"/openpose_img", os.ModePerm)
	os.Mkdir(inputPath+"/openpose_json", os.ModePerm)
	os.Mkdir(inputPath+"/image-parse-v3", os.ModePerm)

	ioutil.WriteFile(inputPath+"/image/test.jpg", image, os.ModePerm)
	ioutil.WriteFile(inputPath+"/cloth/test.jpg", cloth, os.ModePerm)
	ioutil.WriteFile(inputPath+"/cloth-mask/test.jpg", mask, os.ModePerm)
	ioutil.WriteFile(inputPath+"/image-densepose/test.jpg", densepose, os.ModePerm)
	ioutil.WriteFile(inputPath+"/openpose_img/test.png", pose, os.ModePerm)
	ioutil.WriteFile(inputPath+"/openpose_json/test_keypoints.json", []byte(keypoint), os.ModePerm)
	ioutil.WriteFile(inputPath+"/image-parse-v3/test.png", parse, os.ModePerm)
	agnostic, err := Exec(inputPath, outputPath)

	os.RemoveAll(inputPath)
	os.RemoveAll(outputPath)

	return agnostic, err
}

func Exec(inputPath, outputPath string) ([]byte, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	containerConfig := container.Config{
		Image:      "venuns/parse-agnostic",
		Cmd:        []string{"bash", "run.sh"},
		WorkingDir: "/parse_agnostic",
	}

	hostConfig := container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: inputPath,
				Target: "/parse_agnostic/input",
			},
			{
				Type:   mount.TypeBind,
				Source: outputPath,
				Target: "/parse_agnostic/output",
			},
		},
		AutoRemove: true,
	}

	log.Println("parse-agnostic processing")
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

	return ioutil.ReadFile(outputPath + "/test.png")
}
