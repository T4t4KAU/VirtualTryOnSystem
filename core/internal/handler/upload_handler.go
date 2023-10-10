package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"govton/core/internal/logic"
	"govton/core/internal/svc"
	"govton/core/internal/types"
)

func uploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadRequest
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		f, _, err := r.FormFile("file")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		defer f.Close()
		data, err := ioutil.ReadAll(f)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		req.File = data
		req.Path = r.FormValue("path")
		l := logic.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
