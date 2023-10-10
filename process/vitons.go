package process

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"govton/containers/openpose"
	"govton/process/proto"
	"govton/utils"
	"log"
	"time"
)

func VitonsAutoGenerate(image []byte, cloth []byte) []byte {
	mask, err := ProcessClothSegmentation(cloth) // 生成mask图片
	if err != nil {
		logx.Error("colth segmention error:", err)
	}
	logx.Info("cloth segmentation success")

	poseImage, keypoint, err := openpose.Generate(image) // openpose关键点检测
	if err != nil {
		logx.Error("openpose error:", err)
	}
	logx.Info("openpose success")

	densepose, err := ProcessDensePose(image) // 获取densepose图片
	if err != nil {
		logx.Error("densepose error: ", err)
	}
	logx.Info("densepose success")

	parse, _, err := ProcessHumanParse(image) // human parse解析
	if err != nil {
		logx.Error("human parsing error: ", err)
	}
	logx.Info("human parsing success")

	agnostic, err := ProcessParseAgnostic(cloth, mask, image, parse, poseImage, keypoint, densepose) // parsing agnostic
	if err != nil {
		logx.Error("parsing agnostic error: ", err)
	}
	logx.Info("parsing agnostic success")
	result, err := VitonsGenerate(cloth, mask, image, parse, poseImage, string(keypoint), agnostic, densepose)
	if err != nil {
		logx.Error("vitons generate error: ", err)
	}
	logx.Info("vitons generate success")

	return result
}

func VitonsGenerate(cloth []byte, mask []byte, image []byte, parse []byte,
	poseImage []byte, keypoint string, agnostic []byte, densepose []byte) ([]byte, error) {
	conn, err := grpc.Dial(VitonsAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := proto.NewVitonsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()

	log.Println("vitons processing")
	resp, err := client.Generate(ctx, &proto.VitonsRequest{
		Cloth: cloth, Mask: mask, Image: image,
		ImageParse: parse, PoseJson: keypoint,
		ImagePose: poseImage, ParseAgnostic: agnostic,
		ImageDensepose: densepose,
	})
	if err != nil {
		return nil, err
	}
	return utils.ConvertImage(resp.Message), nil
}

func DockerVitonsGenerate(cloth []byte, mask []byte, image []byte, parse []byte, poseImage []byte,
	keypoint string, agnostic []byte, densepose []byte, address string) ([]byte, error) {
	conn, err := grpc.Dial(address+":50059", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := proto.NewVitonsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()

	log.Println("vitons processing")
	resp, err := client.Generate(ctx, &proto.VitonsRequest{
		Cloth: cloth, Mask: mask, Image: image,
		ImageParse: parse, PoseJson: keypoint,
		ImagePose: poseImage, ParseAgnostic: agnostic,
		ImageDensepose: densepose,
	})
	if err != nil {
		return nil, err
	}
	return utils.ConvertImage(resp.Message), nil
}
