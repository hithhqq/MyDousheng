package video

import (
	"fmt"
	"net/http"

	"MyDouSheng/app/video/cmd/api/internal/logic/video"
	"MyDouSheng/app/video/cmd/api/internal/svc"
	"MyDouSheng/app/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// publishvideo
func PublishvideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Printf("err is %v\n", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := video.NewPublishvideoLogic(r.Context(), svcCtx)
		resp, err := l.Publishvideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
