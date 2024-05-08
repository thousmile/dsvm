package logic

import (
	"context"

	"dsvm/internal/svc"
	"dsvm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSchemasLogic struct {
	logx.Logger
	ctx context.Context
	*svc.ServiceContext
}

// NewGetSchemasLogic 获取Schema名称
func NewGetSchemasLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSchemasLogic {
	return &GetSchemasLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		ServiceContext: svcCtx,
	}
}

func (l *GetSchemasLogic) GetSchemas(req *types.GetSchemasRequest) (resp *types.GetSchemasResponse, err error) {
	// todo: add your logic here and delete this line
	return
}
