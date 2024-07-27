package verify

import (
	"fmt"
	"net/http"

	"MyDouSheng/app/identity/cmd/api/internal/logic/verify"
	"MyDouSheng/app/identity/cmd/api/internal/svc"
	"MyDouSheng/common/result"
	"MyDouSheng/common/xerr"

	"github.com/pkg/errors"
)

var ErrTokenExpireError = xerr.NewErrCode(xerr.TOKEN_EXPIRE_ERROR)

// 验证求情token
func VerifyTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var req types.VerifyTokenReq
		// if err := httpx.Parse(r, &req); err != nil {
		// 	fmt.Printf("err is %v\n", err)
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// 	return
		// }

		l := verify.NewVerifyTokenLogic(r.Context(), svcCtx)
		resp, err := l.VerifyToken(r)
		if err == nil && (resp == nil || resp.StatusCode == -1) {
			err = errors.Wrapf(ErrTokenExpireError, "jwtAuthHandler JWT Auth no err , userId is zero ,resp:%+v", resp)
		}
		XUser := "0"
		if resp != nil {
			XUser = fmt.Sprintf("%d", resp.StatusCode)
		}
		w.Header().Set("x-user", XUser)
		result.AuthHttpResult(r, w, resp, err)
	}
}
