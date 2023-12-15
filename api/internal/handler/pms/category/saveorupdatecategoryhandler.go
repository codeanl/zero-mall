package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"simple_mall_new/api/internal/logic/pms/category"
	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"
)

func SaveOrUpdateCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveOrUpdateCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewSaveOrUpdateCategoryLogic(r.Context(), svcCtx)
		resp, err := l.SaveOrUpdateCategory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
