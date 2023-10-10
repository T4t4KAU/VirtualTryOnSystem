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

func ProcessDensePose(image []byte) ([]byte, error) {
	conn, err := grpc.Dial(DensePoseAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := proto.NewDenseposeClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()
	resp, err := client.Generate(ctx, &proto.DensePoseRequest{Image: image})
	if err != nil {
		return nil, err
	}
	return utils.ConvertImage(resp.Message), nil
}

func DockerProcessDensePose(image []byte, address string) ([]byte, error) {
	if address == "" {
		return nil, errors.New("invalid address")
	}
	conn, err := grpc.Dial(address+":50052", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := proto.NewDenseposeClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()

	log.Println("densepose processing")
	resp, err := client.Generate(ctx, &proto.DensePoseRequest{Image: image})
	if err != nil {
		return nil, err
	}
	return utils.ConvertImage(resp.Message), nil
}
