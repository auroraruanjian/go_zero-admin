package user

import (
	"go-zero-demo/api/common/response"
	"net/http"

	"go-zero-demo/api/internal/logic/sys/user"
	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDelLogic(r.Context(), svcCtx)
		resp, err := l.Del(&req)
		response.Response(w, resp, err)
	}
}
