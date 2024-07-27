package video

import (
	"net/http"

	"MyDouSheng/app/video/cmd/api/internal/logic/video"
	"MyDouSheng/app/video/cmd/api/internal/svc"
	"MyDouSheng/app/video/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// getfeed
func GetFeedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFeedReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewGetFeedLogic(r.Context(), svcCtx)
		resp, err := l.GetFeed(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
