package logic

import (
	"context"

	"dsvm/internal/svc"
	"dsvm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitSchemasLogic struct {
	logx.Logger
	ctx context.Context
	*svc.ServiceContext
}

// NewInitSchemasLogic 需要初始化的Schema名称
func NewInitSchemasLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitSchemasLogic {
	return &InitSchemasLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		ServiceContext: svcCtx,
	}
}

func (l *InitSchemasLogic) InitSchemas(req *types.InitSchemasRequest) (resp *types.InitSchemasResponse, err error) {
	// todo: add your logic here and delete this line
	return
}
