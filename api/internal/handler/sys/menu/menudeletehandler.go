package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"simple_mall_new/api/internal/logic/sys/menu"
	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"
)

func MenuDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewMenuDeleteLogic(r.Context(), svcCtx)
		resp, err := l.MenuDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
