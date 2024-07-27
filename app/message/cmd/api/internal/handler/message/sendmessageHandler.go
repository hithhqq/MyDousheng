package message

import (
	"net/http"

	"MyDouSheng/app/message/cmd/api/internal/logic/message"
	"MyDouSheng/app/message/cmd/api/internal/svc"
	"MyDouSheng/app/message/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// sendmessage
func SendmessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendMessageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := message.NewSendmessageLogic(r.Context(), svcCtx)
		resp, err := l.Sendmessage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
