package handler

import (
	"net/http"

	"dsvm/internal/logic"
	"dsvm/internal/svc"
	"dsvm/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// GetSchemasHandler 获取Schema名称
func GetSchemasHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSchemasRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetSchemasLogic(r.Context(), svcCtx)
		resp, err := l.GetSchemas(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
