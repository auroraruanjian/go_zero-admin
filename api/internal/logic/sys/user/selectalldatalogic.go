package user

import (
	"context"

	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectAllDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectAllDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectAllDataLogic {
	return &SelectAllDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectAllDataLogic) SelectAllData(req *types.SelectDataReq) (resp *types.SelectDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
