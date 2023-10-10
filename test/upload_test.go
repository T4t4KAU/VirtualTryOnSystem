package test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

func TestUploadImage(t *testing.T) {
	path := "upload_images/test1.jpg"
	uri := "http://127.0.0.1:8888/upload"
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile("file", path)
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(formFile, f)
	if err != nil {
		t.Fatal(err)
	}

	currentDir, _ := os.Getwd()
	_ = writer.WriteField("path", currentDir+"/testdir/test1.jpg")
	if err = writer.Close(); err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
}
