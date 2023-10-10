package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"govton/core/internal/svc"
	"govton/core/internal/types"
)

type VitonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVitonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VitonLogic {
	return &VitonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VitonLogic) Viton(req *types.VitonRequest) (resp *types.VitonResponse, err error) {
	resp = new(types.VitonResponse)
	resp.Image = req.Image
	return
}
