package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"simple_mall_new/api/internal/logic/sys/menu"
	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"
)

func MenuListsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewMenuListsLogic(r.Context(), svcCtx)
		resp, err := l.MenuLists(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
