package logic

import (
	"context"
	"errors"

	"go-zero-demo/rpc/models/adminmodel"
	"go-zero-demo/rpc/models/query"
	"go-zero-demo/rpc/sys/internal/svc"
	"go-zero-demo/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *sysclient.UserAddReq) (*sysclient.UserAddResp, error) {
	// 先检测用户是否存在，存在则返回错误，不存在则新增
	u := query.AdminUser
	_, uerr := u.Select(u.Id).Where(u.Name.Eq(in.Name)).First()
	if !errors.Is(uerr, gorm.ErrRecordNotFound) {
		return nil, status.Error(codes.InvalidArgument, "用户已存在")
	}

	err := u.WithContext(l.ctx).Create(&adminmodel.AdminUser{
		NickName: in.NickName,
		Name:     in.Name,
		Avatar:   in.Avatar,
		Password: in.Password,
		Email:    in.Email,
		Mobile:   in.Mobile,
		Status:   int(in.Status),
	})

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "用户添加失败")
	}

	return &sysclient.UserAddResp{
		Pong: "success",
	}, nil
}
