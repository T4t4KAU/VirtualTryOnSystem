package test

import (
	"bytes"
	"fmt"
	"govton/containers"
	"govton/containers/agnostic"
	"govton/containers/openpose"
	"govton/containers/resize"
	"govton/containers/resolution"
	"govton/containers/vitons"
	"govton/process"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

const (
	viton_input  = "./viton_input"
	viton_output = "./viton_output"
)

func TestVitonRequest(t *testing.T) {
	imagePath := "images/test1.jpg"
	clothPath := "images/test3.jpg"
	uri := "http://127.0.0.1:8888/viton"

	image, err := os.Open(imagePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer image.Close()

	cloth, err := os.Open(clothPath)
	if err != nil {
		t.Fatal(err)
	}
	defer cloth.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	imagePart, err := writer.CreateFormFile("image", "image.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := io.Copy(imagePart, image); err != nil {
		fmt.Println(err)
		return
	}

	clothPart, err := writer.CreateFormFile("cloth", "test.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := io.Copy(clothPart, cloth); err != nil {
		fmt.Println(err)
		return
	}
	writer.Close()

	// 发送 POST 请求
	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestVitonGenerate(t *testing.T) {
	cloth, _ := ioutil.ReadFile(viton_input + "/cloth/test1.jpg")
	mask, _ := ioutil.ReadFile(viton_input + "/cloth-mask/test1.jpg")
	image, _ := ioutil.ReadFile(viton_input + "/image/test1.jpg")
	parse, _ := ioutil.ReadFile(viton_input + "/image-parsing/parse-agnostic.png")
	poseImage, _ := ioutil.ReadFile(viton_input + "/openpose-img/parse-agnostic.png")
	poseJson, _ := ioutil.ReadFile(viton_input + "/openpose-json/test.json")
	data, _ := process.VitonGenerate(cloth, mask, image, parse, poseImage, string(poseJson))
	file, _ := os.OpenFile(viton_output+"/result.jpg", os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file.Close()
	_, _ = file.Write(data)
}

func TestVitonAutoGenerate(t *testing.T) {
	image, _ := ioutil.ReadFile(viton_input + "/image/test1.jpg")
	cloth, _ := ioutil.ReadFile(viton_input + "/cloth/test1.jpg")
	mask, _ := process.ProcessClothSegmentation(cloth)
	t.Logf("clothmask success\n")

	currentDir, _ := os.Getwd()
	inputPath := currentDir + "/viton_input/image"
	outputPath := currentDir + "/openpose_output"
	keypointPath := currentDir + "/openpose_keypoint"
	poseImage, poseKeypoint, _ := openpose.ExecCPU(inputPath, outputPath, keypointPath)
	t.Logf("openpose success\n")
	parse, _, _ := process.ProcessHumanParse(image)
	t.Logf("parsing success\n")
	data, _ := process.VitonGenerate(cloth, mask, image, parse, poseImage, string(poseKeypoint))
	file, _ := os.OpenFile(viton_output+"/result.jpg", os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file.Close()
	_, _ = file.Write(data)
}

func TestVitonsGenerate(t *testing.T) {
	currentDir, _ := os.Getwd()
	vitons_input := currentDir + "/vitons_input"
	vitons_output := currentDir + "/vitons_output"
	cloth, _ := ioutil.ReadFile(vitons_input + "/cloth/test1.jpg")
	mask, _ := ioutil.ReadFile(vitons_input + "/cloth-mask/test1.jpg")
	image, _ := ioutil.ReadFile(vitons_input + "/image/test1.jpg")
	parse, _ := ioutil.ReadFile(vitons_input + "/image-parsing-v3/parse-agnostic.png")
	poseImage, _ := ioutil.ReadFile(vitons_input + "/openpose_img/parse-agnostic.png")
	poseJson, _ := ioutil.ReadFile(vitons_input + "/openpose_json/test_keypoints.json")
	densepose, _ := ioutil.ReadFile(vitons_input + "/image-densepose/test1.jpg")
	agnostic, _ := ioutil.ReadFile(vitons_input + "/image-parsing-agnostic-v3.2/parse-agnostic.png")
	data, err := process.VitonsGenerate(cloth, mask, image, parse,
		poseImage, string(poseJson), agnostic, densepose)
	if err != nil {
		t.Fatal(err)
	}
	file, _ := os.OpenFile(vitons_output+"/result.jpg", os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file.Close()
	_, _ = file.Write(data)
}

func TestVitonsAutoGenerate(t *testing.T) {
	currentDir, _ := os.Getwd()
	vitons_input := currentDir + "/vitons_input"
	vitons_output := currentDir + "/vitons_output"
	image, _ := ioutil.ReadFile(vitons_input + "/image/test1.jpg")
	cloth, _ := ioutil.ReadFile(vitons_input + "/cloth/test1.jpg")
	mask, _ := process.ProcessClothSegmentation(cloth)
	ioutil.WriteFile(vitons_input+"/cloth-mask/test1.jpg", mask, os.ModePerm)
	t.Logf("clothmask success\n")

	inputPath := vitons_input + "/image"
	outputPath := vitons_input + "/openpose_img"
	keypointPath := vitons_input + "/openpose_json"
	poseImage, poseJson, _ := openpose.Exec(inputPath, outputPath, keypointPath)

	t.Logf("openpose success\n")

	parse, _, err := process.ProcessHumanParse(image)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile(vitons_input+"/image-parsing-v3/parse-agnostic.png", parse, os.ModePerm)
	t.Logf("human-parsing success\n")

	densepose, err := process.ProcessDensePose(image)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile(vitons_input+"/image-densepose/test1.jpg", densepose, os.ModePerm)
	ioutil.WriteFile(vitons_input+"/image-parsing-agnostic/test1.jpg", densepose, os.ModePerm)
	agnostic, _ := process.ProcessParseAgnostic(cloth, mask, image, parse, poseImage, poseJson, densepose)
	ioutil.WriteFile(vitons_input+"/image-parsing-agnostic-v3.2/test1.jpg", agnostic, os.ModePerm)

	data, _ := process.VitonsGenerate(cloth, mask, image, parse, poseImage, string(poseJson), agnostic, densepose)
	file, _ := os.OpenFile(vitons_output+"/result.jpg", os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file.Close()
	file.Write(data)
}

func TestDockerVitonsGenerate(t *testing.T) {
	currentDir, _ := os.Getwd()
	vitons_input := currentDir + "/vitons_input"
	cloth, _ := ioutil.ReadFile(vitons_input + "/cloth/test.jpg")
	mask, _ := ioutil.ReadFile(vitons_input + "/cloth-mask/test.jpg")
	image, _ := ioutil.ReadFile(vitons_input + "/image/test.jpg")
	parse, _ := ioutil.ReadFile(vitons_input + "/image-parsing-v3/parse-agnostic.png")
	poseImage, _ := ioutil.ReadFile(vitons_input + "/openpose_img/parse-agnostic.png")
	keypoint, _ := ioutil.ReadFile(vitons_input + "/openpose_json/test_keypoints.json")
	densepose, _ := ioutil.ReadFile(vitons_input + "/image-densepose/test.jpg")
	agnostic, _ := ioutil.ReadFile(vitons_input + "/image-parsing-agnostic-v3.2/parse-agnostic.png")
	c := vitons.New()
	if err := c.Run(); err != nil {
		t.Fatal(err)
	}
	res, err := process.DockerVitonsGenerate(cloth, mask, image, parse,
		poseImage, string(keypoint), densepose, agnostic, c.Address())
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("images/result.jpg", res, os.ModePerm)
	if err = c.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestDockerVitons(t *testing.T) {
	image, err := ioutil.ReadFile("images/test2.jpg")
	if err != nil {
		t.Fatal(err)
	}
	cloth, err := ioutil.ReadFile("images/test3.jpg")
	if err != nil {
		t.Fatal(err)
	}
	densepose, err := ioutil.ReadFile("images/densepose.jpg")
	if err != nil {
		t.Fatal(err)
	}
	pose, err := ioutil.ReadFile("images/openpose.png")

	keypoint, err := ioutil.ReadFile("images/keypoints.json")
	if err != nil {
		t.Fatal(err)
	}

	parse, err := ioutil.ReadFile("images/parse.png")
	if err != nil {
		t.Fatal(err)
	}
	mask, err := ioutil.ReadFile("images/mask.png")
	if err != nil {
		t.Fatal(err)
	}

	agnostic, err := ioutil.ReadFile("images/parse-agnostic.png")
	if err != nil {
		t.Fatal(err)
	}

	res, err := process.DockerVitonsGenerate(cloth, mask, image, parse, pose, string(keypoint),
		agnostic, densepose, "172.17.0.2")
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("images/result.png", res, os.ModePerm)
}

func TestDockerVitonsProcess(t *testing.T) {
	c := vitons.New()
	if err := c.Run(); err != nil {
		t.Fatal(err)
	}
	if err := c.Process(); err != nil {
		t.Fatal(err)
	}

	println("clothmask address:", containers.ClothMaskAddress)
	println("densepose address:", containers.DensePoseAddress)
	println("human-parse address:", containers.HumanParseAddress)
	println("viton address: ", containers.VitonsAddress)
}

func TestDockerAutoProcess(t *testing.T) {
	image, err := ioutil.ReadFile("images/test6.jpg")
	if err != nil {
		t.Fatal(err)
	}
	cloth, err := ioutil.ReadFile("images/test3.jpg")
	if err != nil {
		t.Fatal(err)
	}

	image, cloth, err = resize.Generate(image, cloth)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("images/resize_image.jpg", image, os.ModePerm)
	ioutil.WriteFile("images/resize_cloth.jpg", cloth, os.ModePerm)

	pose, keypoint, err := openpose.Generate(image)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("images/openpose.png", pose, os.ModePerm)
	ioutil.WriteFile("images/keypoints.json", []byte(keypoint), os.ModePerm)

	mask, err := process.DockerProcessClothSegmentation(cloth, "172.17.0.4")
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("images/mask.png", mask, os.ModePerm)

	parse, _, err := process.DockerProcessHumanParse(image, "172.17.0.5")
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("images/parse.png", parse, os.ModePerm)

	densepose, err := process.DockerProcessDensePose(image, "172.17.0.3")
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("images/densepose.png", densepose, os.ModePerm)

	agnostic, err := agnostic.Generate(image, cloth, mask,
		parse, densepose, pose, string(keypoint))
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("images/parse-agnostic.png", agnostic, os.ModePerm)

	result, err := process.DockerVitonsGenerate(cloth, mask, image, parse,
		pose, keypoint, agnostic, densepose, "172.17.0.2")
	ioutil.WriteFile("images/result.png", result, os.ModePerm)
}

func TestDockerVitonsExec(t *testing.T) {
	image, err := ioutil.ReadFile("images/test1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	cloth, err := ioutil.ReadFile("images/test3.jpg")
	if err != nil {
		t.Fatal(err)
	}
	vitonsContainer := vitons.New()
	if err = vitonsContainer.Run(); err != nil {
		t.Fatal(err)
	}
	if err = vitonsContainer.Process(); err != nil {
		t.Fatal(err)
	}
	pose, keypoint, err := openpose.Generate(image)
	if err != nil {
		t.Fatal(err)
	}

	mask, err := process.DockerProcessClothSegmentation(
		cloth, vitonsContainer.MaskContainer.Address())
	if err != nil {
		t.Fatal(err)
	}

	parse, _, err := process.DockerProcessHumanParse(
		image, vitonsContainer.ParseContainer.Address())
	if err != nil {
		t.Fatal(err)
	}

	densepose, err := process.DockerProcessDensePose(
		image, vitonsContainer.PoseContainer.Address())
	if err != nil {
		t.Fatal(err)
	}

	agnostic, err := agnostic.Generate(image, cloth, mask,
		parse, densepose, pose, string(keypoint))
	if err != nil {
		t.Fatal(err)
	}

	result, err := process.DockerVitonsGenerate(cloth, mask, image, parse,
		pose, keypoint, agnostic, densepose, vitonsContainer.Address())
	ioutil.WriteFile("images/result.png", result, os.ModePerm)
}

func TestVitonsCases(t *testing.T) {
	// num_cases := 20
	format := "test_cases/%d"
	for i := 1; i <= 24; i++ {
		println("========================", "test case", i, "==========================")
		testPath := fmt.Sprintf(format, i)
		image, err := ioutil.ReadFile(testPath + "/image.jpg")
		cloth, err := ioutil.ReadFile(testPath + "/cloth.jpg")
		if err != nil {
			t.Fatal(err)
		}

		image, _ = resolution.Generate(image)
		cloth, _ = resolution.Generate(cloth)
		ioutil.WriteFile(testPath+"/image_res.jpg", image, os.ModePerm)
		ioutil.WriteFile(testPath+"/cloth_res.jpg", cloth, os.ModePerm)

		image, cloth, err = resize.Generate(image, cloth)
		if err != nil {
			t.Fatal(err)
		}
		pose, keypoint, err := openpose.Generate(image)
		if err != nil {
			t.Fatal(err)
		}

		ioutil.WriteFile(testPath+"/openpose.png", pose, os.ModePerm)
		ioutil.WriteFile(testPath+"/keypoints.json", []byte(keypoint), os.ModePerm)

		mask, err := process.DockerProcessClothSegmentation(cloth, "172.17.0.16")
		if err != nil {
			t.Fatal(err)
		}
		ioutil.WriteFile(testPath+"/mask.png", mask, os.ModePerm)

		parse, parse_vis, err := process.DockerProcessHumanParse(image, "172.17.0.17")
		if err != nil {
			t.Fatal(err)
		}
		ioutil.WriteFile(testPath+"/parse.png", parse, os.ModePerm)
		ioutil.WriteFile(testPath+"/parse_vis.png", parse_vis, os.ModePerm)

		densepose, err := process.DockerProcessDensePose(image, "172.17.0.15")
		if err != nil {
			t.Fatal(err)
		}
		ioutil.WriteFile(testPath+"/densepose.png", densepose, os.ModePerm)

		agnostic, err := agnostic.Generate(image, cloth, mask,
			parse, densepose, pose, keypoint)
		if err != nil {
			t.Fatal(err)
		}

		result, err := process.DockerVitonsGenerate(cloth, mask, image, parse,
			pose, keypoint, agnostic, densepose, "172.17.0.14")
		if err != nil {
			t.Fatal(err)
		}
		ioutil.WriteFile(testPath+"/result.png", result, os.ModePerm)
	}
}
