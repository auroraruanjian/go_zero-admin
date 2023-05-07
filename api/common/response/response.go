//utils/response/response.go代码

package response

import (
	"go-zero-demo/api/common/errorx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int `json:"code"`

	Message string `json:"msg"`

	Data interface{} `json:"data,omitempty"`
}

//统一封装成功响应值

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		switch e := err.(type) {
		case *errorx.CodeError: //业务输出错误

			body.Code = e.Code

			body.Message = e.Message

			body.Data = e.Data
		//body.Data = e.Data()
		default: //系统未知错误
			body.Code = 1
			body.Message = err.Error()
		}
	} else {
		body.Code = 0
		body.Message = "请求成功!"
		body.Data = resp
	}

	httpx.OkJson(w, body)
}
