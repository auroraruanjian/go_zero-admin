package logic

import (
	"context"

	"go-zero-demo/rpc/models/query"
	"go-zero-demo/rpc/sys/internal/svc"
	"go-zero-demo/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserLogic) DelUser(in *sysclient.UserDelReq) (*sysclient.UserDelResp, error) {
	if in.Id == 1 {
		return &sysclient.UserDelResp{}, status.Error(codes.InvalidArgument, "超管不可删除")
	}

	// 先检测用户是否存在，存在则返回错误，不存在则新增
	u := query.AdminUser

	user, uerr := u.Select(u.Id).Where(u.Id.Eq(int(in.Id))).First()
	if uerr != nil {
		return nil, status.Error(codes.InvalidArgument, "用户不存在")
	}

	res, del_err := u.Delete(user)
	if del_err != nil {
		return nil, status.Error(codes.InvalidArgument, "用户删除失败:"+del_err.Error())
	}
	if res.RowsAffected > 0 {
		return &sysclient.UserDelResp{
			Pong: "success",
		}, nil
	}

	return &sysclient.UserDelResp{}, nil
}
