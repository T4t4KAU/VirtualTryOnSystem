package test

import (
	"govton/process"
	"io/ioutil"
	"os"
	"testing"
)

func TestClothSegmentation(t *testing.T) {
	image, err := ioutil.ReadFile("clothmask_input/test1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	result, err := process.ProcessClothSegmentation(image)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("clothmask_output/test1.jpg", result, os.ModePerm)
}

func TestDensePose(t *testing.T) {
	image, err := ioutil.ReadFile("densepose_input/test1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	result, err := process.ProcessDensePose(image)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("densepose_output/test1.jpg", result, os.ModePerm)
}

func TestHumanParse(t *testing.T) {
	image, err := ioutil.ReadFile("parse_input/test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	result, vis, err := process.ProcessHumanParse(image)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("parse_output/parse-agnostic.png", result, os.ModePerm)
	ioutil.WriteFile("parse_output/test_vis.png", vis, os.ModePerm)
}

func TestHumanAgnostic(t *testing.T) {
	inputPath := "human_agnostic"
	cloth, _ := ioutil.ReadFile(inputPath + "/cloth/test1.jpg")
	mask, _ := ioutil.ReadFile(inputPath + "/cloth-mask/test1.jpg")
	image, _ := ioutil.ReadFile(inputPath + "/image/test1.jpg")
	densepose, _ := ioutil.ReadFile(inputPath + "/image-densepose/test1.jpg")
	parse, _ := ioutil.ReadFile(inputPath + "/image-parsing-v3/parse-agnostic.png")
	openpose_image, _ := ioutil.ReadFile(inputPath + "/openpose_img/parse-agnostic.png")
	openpose_json, _ := ioutil.ReadFile(inputPath + "/openpose_json/test_keypoints.json")
	parse_agnoic, _ := ioutil.ReadFile(inputPath + "/image-parsing-agnostic/parse-agnostic.png")
	result, err := process.ProcessHumanAgnostic(cloth, mask, image, parse,
		openpose_image, string(openpose_json), parse_agnoic, densepose)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile(inputPath+"/output/parse-agnostic.png", result, os.ModePerm)
}

func TestParseAgnostic(t *testing.T) {
	inputPath := "parse_agnostic"
	cloth, _ := ioutil.ReadFile(inputPath + "/cloth/test1.jpg")
	mask, _ := ioutil.ReadFile(inputPath + "/cloth-mask/test1.jpg")
	image, _ := ioutil.ReadFile(inputPath + "/image/test1.jpg")
	densepose, _ := ioutil.ReadFile(inputPath + "/image-densepose/test1.jpg")
	parse, _ := ioutil.ReadFile(inputPath + "/image-parsing-v3/parse-agnostic.png")
	openpose_image, _ := ioutil.ReadFile(inputPath + "/openpose_img/parse-agnostic.png")
	openpose_json, _ := ioutil.ReadFile(inputPath + "/openpose_json/test_keypoints.json")
	result, err := process.ProcessParseAgnostic(cloth, mask, image, parse,
		openpose_image, string(openpose_json), densepose)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile(inputPath+"/output/parse-agnostic.png", result, os.ModePerm)
}
