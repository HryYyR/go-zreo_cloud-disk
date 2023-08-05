package handler

import (
	"cloud_disk/go-zreo_cloud-disk/core/define"
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
		b := make([]byte, fileheader.Size) //新建一个bytes[]保存文件数据
		_, err = file.Read(b)              //写入文件流
		if err != nil {
			logx.Errorf("文件读取失败: %v", err)
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b)) //通过bytes[]生成md5(hash)
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Xorm.Where("hash=?", hash).Get(rp)
		if err != nil {
			logx.Errorf("数据库查询失败: %v", err)
			return
		}
		if has {
			//判断hash在库中是否存在,如果存在直接返回
			httpx.OkJson(w, &types.FileUploadres{Identity: rp.Identity, Ext: rp.Ext, Name: rp.Name})
			return
		}

		// 存储文件
		var uploadpath string
		if fileheader.Size > int64(define.ChunkDeadline) {
			fmt.Println("开始分片上传")
			uploadpath, err = helper.OssPartUpload(r, b, fileheader) //分片上传云存储
			if err != nil {
				logx.Errorf("分片文件存储失败: %v", err)
				return
			}
		} else {
			fmt.Println("开始直接上传")
			uploadpath, err = helper.OssUpload(r) //上传云存储
			if err != nil {
				logx.Errorf("文件存储失败: %v", err)
				return
			}
		}

		// 往logic传递数据,用于存储进数据库
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
