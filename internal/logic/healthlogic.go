package logic

import (
	"context"
	"github.com/google/uuid"

	"dsvm/internal/svc"
	"dsvm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLogic struct {
	logx.Logger
	ctx context.Context
	*svc.ServiceContext
}

// NewHealthLogic 健康检查
func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		ServiceContext: svcCtx,
	}
}

func (l *HealthLogic) Health() (resp *types.HealthResponse, err error) {
	v7, _ := uuid.NewV7()
	resp = &types.HealthResponse{Message: v7.String()}
	return resp, err
}
