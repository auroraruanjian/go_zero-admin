package logic

import (
	"context"

	"go-zero-demo/rpc/common/service"
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
	//user_row, err := u.Where(u.Id.Eq(in.UserId)).First()
	user_row, err := u.Preload(u.AdminRole).Where(u.Id.Eq(int(in.UserId))).First()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "用户不存在")
	}

	// 角色字符串切片
	role_id_list := []int{}
	// 角色对象
	adminRole := []*sysclient.AdminRole{}
	for _, value := range user_row.AdminRole {
		adminRole = append(adminRole, &sysclient.AdminRole{
			Name: value.Name,
			Slug: value.Slug,
		})
		role_id_list = append(role_id_list, int(value.Id))
	}

	// 获取角色权限
	UserService := service.AdminUserService{
		DB: l.svcCtx.DB,
	}
	permission := UserService.GetAdminRolePermission(role_id_list)

	return &sysclient.InfoResp{
		Avatar:          user_row.Avatar,
		Name:            user_row.Name,
		NickName:        user_row.NickName,
		Email:           user_row.Email,
		Mobile:          user_row.Mobile,
		CreatedAt:       user_row.CreatedAt.Unix(),
		AdminRole:       adminRole,
		AdminPermission: permission,
	}, nil
}
