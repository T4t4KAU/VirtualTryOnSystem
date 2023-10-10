package process

import (
	"context"
	"google.golang.org/grpc"
	"govton/process/proto"
	"govton/utils"
	"time"
)

func ProcessHumanAgnostic(cloth []byte, mask []byte, image []byte, parse []byte,
	poseImage []byte, poseJson string, agnostic []byte, densepose []byte) ([]byte, error) {
	conn, err := grpc.Dial(HumanAgnosticAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := proto.NewHumanAgnosticClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()

	resp, err := client.Generate(ctx, &proto.HumanAgnosticRequest{
		Cloth: cloth, Mask: mask, Image: image,
		ImageParse: parse, PoseJson: poseJson,
		ImagePose: poseImage, ParseAgnostic: agnostic,
		ImageDensepose: densepose,
	})
	if err != nil {
		return nil, err
	}
	return utils.ConvertImage(resp.Message), nil
}
