package process

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"govton/process/proto"
	"govton/utils"
	"log"
	"time"
)

func ProcessParseAgnostic(cloth []byte, mask []byte, image []byte, parse []byte,
	poseImage []byte, poseJson string, densepose []byte) ([]byte, error) {
	conn, err := grpc.Dial(ParseAgnosticAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := proto.NewParseAgnosticClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()
	resp, err := client.Generate(ctx, &proto.ParseAgnosticRequest{
		Cloth: cloth, Mask: mask, Image: image,
		ImageParse: parse, PoseJson: poseJson,
		ImagePose:      poseImage,
		ImageDensepose: densepose,
	})
	if err != nil {
		return nil, err
	}
	return utils.ConvertImage(resp.Message), nil
}

func DockerProcessParseAgnostic(cloth []byte, mask []byte, image []byte, parse []byte,
	poseImage []byte, poseJson string, densepose []byte, address string) ([]byte, error) {
	if address == "" {
		return nil, errors.New("invalid address")
	}
	conn, err := grpc.Dial(address+":50055", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := proto.NewParseAgnosticClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()

	log.Println("parse-agnostic processing")
	resp, err := client.Generate(ctx, &proto.ParseAgnosticRequest{
		Cloth: cloth, Mask: mask, Image: image,
		ImageParse: parse, PoseJson: poseJson,
		ImagePose:      poseImage,
		ImageDensepose: densepose,
	})
	if err != nil {
		return nil, err
	}
	return utils.ConvertImage(resp.Message), nil
}
