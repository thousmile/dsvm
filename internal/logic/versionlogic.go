package logic

import (
	"context"
	"fmt"

	"dsvm/internal/svc"
	"dsvm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VersionLogic struct {
	logx.Logger
	ctx context.Context
	*svc.ServiceContext
}

// NewVersionLogic APP 版本
func NewVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VersionLogic {
	return &VersionLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		ServiceContext: svcCtx,
	}
}

func (l *VersionLogic) Version() (resp *types.VersionResponse, err error) {
	resp = &types.VersionResponse{
		Version: fmt.Sprintf("%s -> %s", l.Config.Name, l.Config.Version),
	}
	return resp, err
}
