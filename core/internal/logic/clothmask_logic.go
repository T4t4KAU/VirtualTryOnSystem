package logic

import (
	"context"
	"govton/process"

	"govton/core/internal/svc"
	"govton/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClothmaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClothmaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClothmaskLogic {
	return &ClothmaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClothmaskLogic) Clothmask(req *types.ClothMaskRequest) (resp *types.ClothMaskRequest, err error) {
	image, err := process.ProcessClothSegmentation(req.Image)
	if err != nil {
		return nil, err
	}
	resp.Image = image

	return
}
