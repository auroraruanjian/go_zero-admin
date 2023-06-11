package middleware

import (
	"encoding/json"
	"fmt"
	"go-zero-demo/api/common/errorx"
	"go-zero-demo/api/common/response"
	"go-zero-demo/rpc/sys/sys"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUrlMiddleware struct {
	Sys sys.Sys
}

func NewCheckUrlMiddleware(sys sys.Sys) *CheckUrlMiddleware {
	return &CheckUrlMiddleware{
		Sys: sys,
	}
}

func (m *CheckUrlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := r.Context()
		//l := logx.WithContext(r.Context())
		// 检查用户是否有菜单权限
		user_id, err := c.Value("userId").(json.Number).Int64()

		if err != nil {
			logx.Errorf("缺少必要参数 user-id")
			response.Response(w, nil, errorx.NewDefaultError("对不起，您没有权限"))
		}

		resp, permiss_err := m.Sys.CheckPermission(c, &sys.CheckPermissionReq{
			UserId: int32(user_id),
			Rule:   r.RequestURI,
		})

		if permiss_err != nil {
			logx.Errorf(permiss_err.Error())
			response.Response(w, nil, errorx.NewDefaultError("对不起，您没有权限:"+permiss_err.Error()))
		}

		fmt.Print(resp)
		if resp.Pong == "success" {
			// Passthrough to next handler if need
			next(w, r)
		} else {
			response.Response(w, nil, errorx.NewDefaultError("对不起，未知的错误"))
		}
	}
}
