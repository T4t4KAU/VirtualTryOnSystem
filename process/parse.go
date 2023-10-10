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

func ProcessHumanParse(image []byte) ([]byte, []byte, error) {
	conn, err := grpc.Dial(HumanParseAddress, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	client := proto.NewHumanParseClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*50)
	defer cancel()
	resp, err := client.Generate(ctx, &proto.ParseRequest{Image: image})
	if err != nil {
		return nil, nil, err
	}
	return utils.ConvertImage(resp.Parse), utils.ConvertImage(resp.ParseVis), nil
}

func DockerProcessHumanParse(image []byte, address string) ([]byte, []byte, error) {
	if address == "" {
		return nil, nil, errors.New("invalid address")
	}
	conn, err := grpc.Dial(address+":50053", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	client := proto.NewHumanParseClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*50)
	defer cancel()

	log.Println("human-parse processing")
	resp, err := client.Generate(ctx, &proto.ParseRequest{Image: image})
	if err != nil {
		return nil, nil, err
	}
	return utils.ConvertImage(resp.Parse), utils.ConvertImage(resp.ParseVis), nil
}
