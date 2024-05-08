package handler

import (
	"net/http"

	"dsvm/internal/logic"
	"dsvm/internal/svc"
	"dsvm/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

// InitSchemasHandler 需要初始化的Schema名称
func InitSchemasHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InitSchemasRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewInitSchemasLogic(r.Context(), svcCtx)
		resp, err := l.InitSchemas(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
