package logic

import (
	"context"

	"gozeroapi/internal/svc"
	"gozeroapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GozeroapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGozeroapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GozeroapiLogic {
	return &GozeroapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GozeroapiLogic) Gozeroapi(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
