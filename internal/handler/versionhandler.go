package handler

import (
	"net/http"

	"dsvm/internal/logic"
	"dsvm/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// VersionHandler APP 版本
func VersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewVersionLogic(r.Context(), svcCtx)
		resp, err := l.Version()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
