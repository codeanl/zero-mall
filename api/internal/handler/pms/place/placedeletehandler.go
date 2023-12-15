package place

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"simple_mall_new/api/internal/logic/pms/place"
	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"
)

func PlaceDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeletePlaceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := place.NewPlaceDeleteLogic(r.Context(), svcCtx)
		resp, err := l.PlaceDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
