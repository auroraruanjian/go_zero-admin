package user

import (
	"go-zero-demo/api/common/response"
	"net/http"

	"go-zero-demo/api/internal/logic/sys/user"
	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewAddLogic(r.Context(), svcCtx)
		resp, err := l.Add(&req)
		response.Response(w, resp, err)
	}
}
