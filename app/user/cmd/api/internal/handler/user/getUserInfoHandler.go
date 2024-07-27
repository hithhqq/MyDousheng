package user

import (
	"fmt"
	"net/http"

	"MyDouSheng/app/user/cmd/api/internal/logic/user"
	"MyDouSheng/app/user/cmd/api/internal/svc"
	"MyDouSheng/app/user/cmd/api/internal/types"
	middleware "MyDouSheng/common/globalmiddleware"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// userinfo
func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Println("claims", r.Context().Value("claims"))
		if r.Context().Value(middleware.ClaimsKey) != nil {

			httpx.OkJsonCtx(r.Context(), w, &types.UserInfoRsp{
				StatusCode: -1,
				StatusMsg:  "登录失效，请重新登录",
			})
			return
		}
		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
