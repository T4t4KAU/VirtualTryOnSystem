package logic

import (
	"context"
	"govton/process"

	"govton/core/internal/svc"
	"govton/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ParseAgnosticLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewParseAgnosticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ParseAgnosticLogic {
	return &ParseAgnosticLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ParseAgnosticLogic) ParseAgnostic(req *types.ParseAgnosticRequest) (resp *types.ParseAgnosticResponse, err error) {
	agnostic, err := process.ProcessParseAgnostic(req.Cloth, req.Mask, req.Image,
		req.ImageParse, req.OpenPoseImage, req.OpenPoseKeypoint, req.DensePoseImage)
	resp.Image = agnostic
	return
}
