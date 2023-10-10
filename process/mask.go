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

func ProcessClothSegmentation(image []byte) ([]byte, error) {
	conn, err := grpc.Dial(ClothMaskAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := proto.NewClothSegmentationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()
	resp, err := client.Generate(ctx, &proto.ImageRequest{Image: image})
	if err != nil {
		return nil, err
	}
	return utils.ConvertImage(resp.Message), nil
}

func DockerProcessClothSegmentation(image []byte, address string) ([]byte, error) {
	if address == "" {
		return nil, errors.New("invalid address")
	}
	conn, err := grpc.Dial(address+":50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	log.Println("cloth segmentation processing")
	client := proto.NewClothSegmentationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()
	resp, err := client.Generate(ctx, &proto.ImageRequest{Image: image})
	if err != nil {
		return nil, err
	}
	return utils.ConvertImage(resp.Message), nil
}
