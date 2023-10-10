package handler

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"govton/core/internal/logic"
	"govton/core/internal/svc"
	"govton/core/internal/types"
)

func vitonHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var req types.VitonRequest
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		image, _, err := r.FormFile("image")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		defer image.Close()
		cloth, _, err := r.FormFile("cloth")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		defer cloth.Close()
		imageBytes, _ := ioutil.ReadAll(image)
		clothBytes, _ := ioutil.ReadAll(cloth)
		req.Image = imageBytes
		req.Cloth = clothBytes
		l := logic.NewVitonLogic(r.Context(), svcCtx)
		resp, err := l.Viton(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.Header().Set("Content-Type", "image/png")
			w.Header().Set("Content-Length", strconv.Itoa(len(resp.Image)))
			w.Write(resp.Image)
		}
	}
}
