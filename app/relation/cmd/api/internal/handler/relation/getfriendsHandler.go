package relation

import (
	"net/http"

	"MyDouSheng/app/relation/cmd/api/internal/logic/relation"
	"MyDouSheng/app/relation/cmd/api/internal/svc"
	"MyDouSheng/app/relation/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// getfriends
func GetfriendsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFriendsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := relation.NewGetfriendsLogic(r.Context(), svcCtx)
		resp, err := l.Getfriends(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
