package logic

import (
	"context"
	"govton/containers/densepose"
	"govton/core/internal/svc"
	"govton/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DenseposeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDenseposeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DenseposeLogic {
	return &DenseposeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DenseposeLogic) Densepose(req *types.DensePoseRequest) (resp *types.DensePoseResponse, err error) {
	c := densepose.New()
	err = c.Run()
	if err != nil {
		return nil, err
	}
	result, err := c.Generate()
	if err != nil {
		return nil, err
	}

	resp = new(types.DensePoseResponse)
	resp.Image = result

	err = c.Stop()
	if err != nil {
		return resp, err
	}
	return
}
