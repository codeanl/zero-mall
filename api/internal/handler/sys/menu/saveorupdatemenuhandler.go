package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"simple_mall_new/api/internal/logic/sys/menu"
	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"
)

func SaveOrUpdateMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveOrUpdateMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewSaveOrUpdateMenuLogic(r.Context(), svcCtx)
		resp, err := l.SaveOrUpdateMenu(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
