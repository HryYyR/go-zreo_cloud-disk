package handler

import (
	"cloud_disk/go-zreo_cloud-disk/core/internal/logic"
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRepositorySaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRepositorySavereq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserRepositorySaveLogic(r.Context(), svcCtx)
		resp, err := l.UserRepositorySave(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			fmt.Println("Parse2error:", err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
