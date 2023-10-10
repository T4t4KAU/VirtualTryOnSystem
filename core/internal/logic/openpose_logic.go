package logic

import (
	"context"
	"govton/containers/openpose"

	"govton/core/internal/svc"
	"govton/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenposeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpenposeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenposeLogic {
	return &OpenposeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpenposeLogic) Openpose(req *types.OpenPoseRequest) (resp *types.OpenPoseResponse, err error) {
	image, keypoint, err := openpose.ExecCPU(req.InputPath, req.OutputPath, req.KeypointPath)
	if err != nil {
		return nil, err
	}

	resp = new(types.OpenPoseResponse)
	resp.Image = image
	resp.Keypoint = keypoint

	return
}
