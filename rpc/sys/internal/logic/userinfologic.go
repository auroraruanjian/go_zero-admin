package logic

import (
	"context"
	"fmt"

	"go-zero-demo/rpc/models/query"
	"go-zero-demo/rpc/sys/internal/svc"
	"go-zero-demo/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *sysclient.InfoReq) (*sysclient.InfoResp, error) {
	// todo: add your logic here and delete this line
	//in.UserId
	u := query.AdminUser
	user_row, err := u.Where(u.Id.Eq(in.UserId)).First()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "用户不存在")
	}
	fmt.Println(user_row)

	return &sysclient.InfoResp{
		Avatar:    user_row.Avatar,
		Name:      user_row.Name,
		NickName:  user_row.NickName,
		Email:     user_row.Email,
		Mobile:    user_row.Mobile,
		CreatedAt: user_row.CreatedAt.Unix(),
	}, nil
}
