package process

import (
	"context"
	"encoding/base64"
	"google.golang.org/grpc"
	"govton/process/proto"
	"log"
	"time"
)

const (
	VitonAddress = "172.17.0.2:50050"
)

func VitonGenerate(cloth, mask, image, parse, poseImage []byte, keypoint string) ([]byte, string) {
	conn, err := grpc.Dial(VitonAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error: %v", err.Error())
	}
	defer conn.Close()
	client := proto.NewVitonClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()
	res, err := client.Generate(ctx, &proto.DataRequest{Name: "vitons", Cloth: cloth,
		ClothMask: mask, Image: image, ImageParse: parse,
		ImagePose: poseImage, PoseJson: string(keypoint)})
	if err != nil {
		log.Fatalf("send error: %v", err.Error())
	}
	data, _ := base64.StdEncoding.DecodeString(string(res.Result))
	return data, string(res.Result)
}

func DockerVitonGenerate(cloth, mask, image, parse, poseImage []byte, keypoint, address string) ([]byte, string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error: %v", err.Error())
	}
	defer conn.Close()
	client := proto.NewVitonClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()
	res, err := client.Generate(ctx, &proto.DataRequest{Name: "vitons", Cloth: cloth,
		ClothMask: mask, Image: image, ImageParse: parse,
		ImagePose: poseImage, PoseJson: string(keypoint)})
	if err != nil {
		log.Fatalf("send error: %v", err.Error())
	}
	data, _ := base64.StdEncoding.DecodeString(string(res.Result))
	return data, string(res.Result)
}
