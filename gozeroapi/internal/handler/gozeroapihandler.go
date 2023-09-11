package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozeroapi/internal/logic"
	"gozeroapi/internal/svc"
	"gozeroapi/internal/types"
)

func GozeroapiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGozeroapiLogic(r.Context(), svcCtx)
		resp, err := l.Gozeroapi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
