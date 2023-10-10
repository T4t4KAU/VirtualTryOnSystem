package logic

import (
	"context"
	"io/ioutil"
	"os"

	"govton/core/internal/svc"
	"govton/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadRequest) (resp *types.UploadResponse, err error) {
	err = ioutil.WriteFile(req.Path, req.File, os.ModePerm)
	if err != nil {
		return
	}
	resp.Path = req.Path
	return
}
