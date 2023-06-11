package logic

import (
	"context"
	"database/sql"

	"go-zero-demo/rpc/sys/internal/svc"
	"go-zero-demo/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckPermissionLogic {
	return &CheckPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckPermissionLogic) CheckPermission(in *sysclient.CheckPermissionReq) (*sysclient.CheckPermissionResp, error) {
	var permiss []int64

	res := l.svcCtx.DB.Raw(`
			SELECT 
				admin_permission_id 
			FROM admin_role_permission
			WHERE admin_permission_id IN(
				SELECT id FROM admin_permissions WHERE rule = @rule
			)
			AND admin_role_id IN(
				SELECT admin_user_id from admin_role_users WHERE admin_user_id= @user_id
			)
		`, sql.Named("rule", in.Rule),
		sql.Named("user_id", in.UserId),
	).Find(&permiss)

	if res.Error != nil {
		return nil, status.Error(codes.Aborted, "未知的错误，"+res.Error.Error())
	}

	if len(permiss) > 0 {
		return &sysclient.CheckPermissionResp{
			Pong: "success",
		}, nil
	}

	return &sysclient.CheckPermissionResp{}, nil
}
