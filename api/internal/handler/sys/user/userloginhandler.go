package user

import (
	"net/http"

	"go-zero-demo/api/common/response"
	"go-zero-demo/api/internal/logic/sys/user"
	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		response.Response(w, resp, err)
	}
}
