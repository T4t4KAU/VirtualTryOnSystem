package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"govton/core/internal/logic"
	"govton/core/internal/svc"
	"govton/core/internal/types"
)

func openposeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OpenPoseRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOpenposeLogic(r.Context(), svcCtx)
		resp, err := l.Openpose(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
