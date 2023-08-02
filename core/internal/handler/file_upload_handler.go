package handler

import (
	"cloud_disk/go-zreo_cloud-disk/core/helper"
	"cloud_disk/go-zreo_cloud-disk/core/internal/logic"
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"cloud_disk/go-zreo_cloud-disk/core/models"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadreq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, fileheader, err := r.FormFile("file")
		if err != nil {
			logx.Errorf("文件不存在: %v", err)
			return
		}

		// 判断文件是否存在
		b := make([]byte, fileheader.Size)
		_, err = file.Read(b)
		if err != nil {
			logx.Errorf("文件读取失败: %v", err)
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Xorm.Where("hash=?", hash).Get(rp)
		if err != nil {
			logx.Errorf("数据库查询失败: %v", err)
			return
		}
		if has {
			httpx.OkJson(w, &types.FileUploadres{Identity: rp.Identity, Ext: rp.Ext, Name: rp.Name})
			return
		}

		// 存储文件
		uploadpath, err := helper.OssUpload(r)
		if err != nil {
			logx.Errorf("文件存储失败: %v", err)
			return
		}

		// 往logic传递数据
		req.Name = fileheader.Filename
		req.Ext = path.Ext(fileheader.Filename)
		req.Size = fileheader.Size
		req.Hash = hash
		req.Path = uploadpath
		// end
		logx.Debug("ok")

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
