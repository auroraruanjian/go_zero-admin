package user

import (
	"context"
	"encoding/json"

	"go-zero-demo/api/common/errorx"
	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"
	"go-zero-demo/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResp, err error) {
	user_id, _ := l.ctx.Value("userId").(json.Number).Int64()

	info_resp, info_err := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{
		UserId: user_id,
	})

	if info_err != nil {
		fromError, ok := status.FromError(info_err)
		if !ok {
			l.Logger.Errorf("RPC异常:%s", info_err.Error())
		}

		l.Logger.Errorf("根据用户ID: %s 查询错误: %s", user_id, info_err.Error())
		return nil, errorx.NewDefaultError(fromError.Message())
	}

	return &types.UserInfoResp{
		Code:      "000000",
		Message:   "登录成功",
		Avatar:    info_resp.Avatar,
		Name:      info_resp.Name,
		NickName:  info_resp.NickName,
		Email:     info_resp.Email,
		Mobile:    info_resp.Mobile,
		CreatedAt: info_resp.CreatedAt,
	}, nil
}
