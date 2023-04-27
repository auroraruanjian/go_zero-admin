package user

import (
	"context"
	"encoding/json"
	"strings"

	"go-zero-demo/api/common/errorx"
	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"
	"go-zero-demo/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		resStr, _ := json.Marshal(req)
		l.Logger.Errorf("账户或密码为空，请求：%s", resStr)
		return nil, errorx.NewDefaultError("账户或密码错误！")
	}

	login_resp, login_err := l.svcCtx.Sys.Login(l.ctx, &sysclient.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})

	if login_err != nil {
		fromError, ok := status.FromError(login_err)
		if !ok {
			l.Logger.Errorf("RPC异常:%s", login_err.Error())
		}
		// 判断服务端返回的是否是指定code的错误
		/*
			if fromError.Code() == codes.InvalidArgument {
				l.Logger.Errorf("invalid arguments")
			}
		*/
		l.Logger.Errorf("根据用户名: %s和密码: %s查询用户异常:%s", req.UserName, req.Password, login_err.Error())
		return nil, errorx.NewDefaultError(fromError.Message())
	}

	return &types.LoginResp{
		Code:         "000000",
		Message:      "登录成功",
		Status:       login_resp.Status,
		UserName:     login_resp.UserName,
		AccessToken:  login_resp.AccessToken,
		AccessExpire: login_resp.AccessExpire,
		RefreshAfter: login_resp.RefreshAfter,
	}, nil
}
