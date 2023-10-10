package logic

import (
	"context"
	"govton/containers"
	"govton/core/internal/svc"
	"govton/core/internal/types"
	"govton/process"

	"github.com/zeromicro/go-zero/core/logx"
)

type ParseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewParseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ParseLogic {
	return &ParseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ParseLogic) Parse(req *types.HumanParseRequest) (resp *types.HumanParseResponse, err error) {
	parse, _, err := process.DockerProcessHumanParse(req.Image, containers.HumanParseAddress)

	resp = new(types.HumanParseResponse)
	resp.Image = parse
	return
}
