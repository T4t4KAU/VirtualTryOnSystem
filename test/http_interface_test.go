package test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

func TestOpenPoseRequest(t *testing.T) {
	// 读取图片数据
	file, err := os.Open("./openpose_input/test1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// 创建一个 multipart/form-data 请求体
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", "test1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}
	writer.Close()

	// 发送 POST 请求
	req, err := http.NewRequest("POST", "http://localhost:8888/openpose", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
}
