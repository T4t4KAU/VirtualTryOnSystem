package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"govton/containers"
	"govton/containers/agnostic"
	"govton/containers/openpose"
	"govton/containers/resize"
	"govton/containers/resolution"
	"govton/core/internal/svc"
	"govton/core/internal/types"
	"govton/process"
	"log"
	"sync"
)

var (
	index int         = 0
	mutex *sync.Mutex = new(sync.Mutex)
)

type VitonsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVitonsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VitonsLogic {
	return &VitonsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VitonsLogic) Vitons(req *types.VitonsRequest) (resp *types.VitonsResponse, err error) {
	resp = new(types.VitonsResponse)
	resp.Result, err = Generate(req.Image, req.Cloth)
	return
}

// 执行处理流程
func Generate(image, cloth []byte) ([]byte, error) {
	mutex.Lock()
	defer mutex.Unlock()
	index++
	fmt.Printf("=================== process %d ===================\n", index)

	image, _ = resolution.Generate(image)
	cloth, _ = resolution.Generate(cloth)
	image, cloth, err := resize.Generate(image, cloth)
	if err != nil {
		return nil, err
	}
	pose, keypoint, err := openpose.Generate(image)
	if err != nil {
		return nil, err
	}

	mask, err := process.DockerProcessClothSegmentation(cloth, containers.ClothMaskAddress)
	if err != nil {
		return nil, err
	}

	parse, _, err := process.DockerProcessHumanParse(image, containers.HumanParseAddress)
	if err != nil {
		return nil, err
	}

	densepose, err := process.DockerProcessDensePose(image, containers.DensePoseAddress)
	if err != nil {
		return nil, err
	}

	agnostic, err := agnostic.Generate(image, cloth, mask,
		parse, densepose, pose, string(keypoint))
	if err != nil {
		return nil, err
	}

	viton, err := process.DockerVitonsGenerate(cloth, mask, image, parse,
		pose, keypoint, agnostic, densepose, containers.VitonsAddress)
	if err != nil {
		return nil, err
	}

	log.Println("viton finish")
	return viton, err
}
