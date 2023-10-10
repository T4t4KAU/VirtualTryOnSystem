package test

import (
	"govton/containers/densepose"
	"govton/containers/mask"
	"govton/containers/parsing"
	"govton/containers/resize"
	"govton/containers/resolution"
	"govton/process"
	"io/ioutil"
	"os"
	"testing"
)

func TestDockerHumanParseRun(t *testing.T) {
	c := parsing.New()
	if err := c.Run(); err != nil {
		t.Fatal(err)
	}
	println(c.Address())
	if err := c.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestHumanParseGenerate(t *testing.T) {
	image, err := ioutil.ReadFile("images/test.jpg")
	if err != nil {
		t.Fatal(err)
	}

	c := parsing.New()
	if err = c.Run(); err != nil {
		t.Fatal()
	}

	res, vis, err := process.DockerProcessHumanParse(image, c.Address())
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("images/parse.png", res, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("images/parse_vis.png", vis, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	if err = c.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestDockerDensePoseRun(t *testing.T) {
	c := densepose.New()
	if err := c.Run(); err != nil {
		t.Fatal(err)
	}
	println(c.Address())
	if err := c.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestDensePoseGenerate(t *testing.T) {
	image, err := ioutil.ReadFile("densepose_input/test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	c := densepose.New()
	if err = c.Run(); err != nil {
		t.Fatal(err)
	}
	res, err := process.DockerProcessDensePose(image, c.Address())
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("densepose_output/test.png", res, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	if err = c.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestDockerClothMaskGenerate(t *testing.T) {
	cloth, _ := ioutil.ReadFile("./clothmask_input/test.jpg")
	c := mask.New()
	if err := c.Run(); err != nil {
		t.Fatal(err)
	}
	res, err := process.DockerProcessClothSegmentation(cloth, c.Address())
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("./clothmask_output/test.jpg", res, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDockerResizeGenerate(t *testing.T) {
	image, err := ioutil.ReadFile("images/test2.jpg")
	if err != nil {
		t.Fatal(err)
	}
	cloth, err := ioutil.ReadFile("images/test3.jpg")
	if err != nil {
		t.Fatal(err)
	}
	resizeImage, resizeCloth, err := resize.Generate(image, cloth)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("images/resize_image.jpg", resizeImage, os.ModePerm)
	ioutil.WriteFile("images/resize_cloth.jpg", resizeCloth, os.ModePerm)
}

func TestResolutionGenerate(t *testing.T) {
	image, err := ioutil.ReadFile("resolution_input/test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	res, err := resolution.Generate(image)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("resolution_output/test.jpg", res, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}
