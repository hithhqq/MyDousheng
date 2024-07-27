package middleware

import (
	"MyDouSheng/app/user/cmd/rpc/userservice"
	"MyDouSheng/common/globalmiddleware"
	"encoding/json"
	"fmt"
	"net/http"
)

type JwtauthMiddleware struct {
}

func NewJwtauthMiddleware() *JwtauthMiddleware {
	return &JwtauthMiddleware{}
}

func (m *JwtauthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		if r.Context().Value(globalmiddleware.ClaimsKey) == "-1" {
			response := &userservice.LoginResp{
				StatusCode: -1,
				StatusMsg:  "登录失效, 请重新登录",
			}
			jsonData, err := json.Marshal(response)
			if err != nil {
				// 如果序列化失败，返回错误信息
				fmt.Printf("err:%v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(jsonData)
			return
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}
