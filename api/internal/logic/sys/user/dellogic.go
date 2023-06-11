package user

import (
	"context"

	"go-zero-demo/api/common/errorx"
	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"
	"go-zero-demo/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.DelUserReq) (resp *types.DelUserResp, err error) {
	UserResp, err := l.svcCtx.Sys.DelUser(l.ctx, &sys.UserDelReq{
		Id: req.Id,
	})

	if err != nil {
		fromError, ok := status.FromError(err)
		if !ok {
			l.Logger.Errorf("RPC异常:%s", err.Error())
		}

		l.Logger.Errorf("用户ID: %d 删除失败: %s", req.Id, err.Error())
		return nil, errorx.NewDefaultError(fromError.Message())
	}

	if UserResp.Pong != "success" {
		l.Logger.Errorf("用户ID: %d 删除失败: %s", req.Id, UserResp.String())
		return nil, errorx.NewDefaultError("未知的异常")
	}

	return &types.DelUserResp{}, nil
}
