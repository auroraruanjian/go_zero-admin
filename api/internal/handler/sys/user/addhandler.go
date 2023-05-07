package user

import (
	"go-zero-demo/api/common/response"
	"net/http"

	"go-zero-demo/api/internal/logic/sys/user"
	"go-zero-demo/api/internal/svc"
)

func AddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewAddLogic(r.Context(), svcCtx)
		resp, err := l.Add()
		response.Response(w, resp, err)
	}
}
