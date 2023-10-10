package logic

import (
	"context"
	"govton/process"

	"govton/core/internal/svc"
	"govton/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HumanAgnosticLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHumanAgnosticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HumanAgnosticLogic {
	return &HumanAgnosticLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HumanAgnosticLogic) HumanAgnostic(req *types.HumanAgnosticRequest) (resp *types.HumanAgnosticResponse, err error) {
	agnostic, err := process.ProcessParseAgnostic(req.Cloth, req.Mask, req.Image,
		req.ImageParse, req.OpenPoseImage, req.OpenPoseKeypoint, req.DensePoseImage)
	if err != nil {
		return nil, err
	}
	resp.Image = agnostic
	return
}
