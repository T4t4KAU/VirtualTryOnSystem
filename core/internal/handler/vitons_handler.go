package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"govton/core/internal/logic"
	"govton/core/internal/svc"
	"govton/core/internal/types"
	"io/ioutil"
	"net/http"
	"strconv"
)

func vitonsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var req types.VitonsRequest
		_ = r.ParseMultipartForm(32 << 20)

		image, _, err := r.FormFile("image")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		defer image.Close()
		imageBytes, err := ioutil.ReadAll(image)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		cloth, _, err := r.FormFile("cloth")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		defer cloth.Close()
		clothBytes, err := ioutil.ReadAll(cloth)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		req.Image = imageBytes
		req.Cloth = clothBytes

		l := logic.NewVitonsLogic(r.Context(), svcCtx)
		resp, err := l.Vitons(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.Header().Set("Content-Type", "image/png")
			w.Header().Set("Content-Length", strconv.Itoa(len(resp.Result)))
			w.Write(resp.Result)
		}
	}
}
