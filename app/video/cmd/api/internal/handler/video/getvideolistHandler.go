package video

import (
	"net/http"

	"MyDouSheng/app/video/cmd/api/internal/logic/video"
	"MyDouSheng/app/video/cmd/api/internal/svc"
	"MyDouSheng/app/video/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// getvideolist
func GetvideolistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetVideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewGetvideolistLogic(r.Context(), svcCtx)
		resp, err := l.Getvideolist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
