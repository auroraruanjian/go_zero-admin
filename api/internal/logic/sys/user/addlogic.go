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

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddUserReq) (resp *types.AddUserResp, err error) {
	//admin_id, _ := l.ctx.Value("userId").(json.Number).Int64()
	//fmt.Print(admin_id)

	UserResp, err := l.svcCtx.Sys.AddUser(l.ctx, &sys.UserAddReq{
		Name:     req.Name,
		NickName: req.NickName,
		Avatar:   req.Avatar,
		Password: req.Password,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Status:   req.Status,
	})

	if err != nil {
		fromError, ok := status.FromError(err)
		if !ok {
			l.Logger.Errorf("RPC异常:%s", err.Error())
		}

		l.Logger.Errorf("用户名: %s 添加失败: %s", req.Name, err.Error())
		return nil, errorx.NewDefaultError(fromError.Message())
	}

	if UserResp.Pong != "success" {
		l.Logger.Errorf("用户名: %s 添加失败: %s", req.Name, UserResp.String())
		return nil, errorx.NewDefaultError("未知的异常")
	}

	return &types.AddUserResp{}, nil
}
