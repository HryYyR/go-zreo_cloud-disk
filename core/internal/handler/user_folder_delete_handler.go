package handler

import (
	"cloud_disk/go-zreo_cloud-disk/core/internal/logic"
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFolderDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFolderDeletereq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserFolderDeleteLogic(r.Context(), svcCtx)
		resp, err := l.UserFolderDelete(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
