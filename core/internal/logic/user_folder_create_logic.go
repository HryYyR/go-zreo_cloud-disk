package logic

import (
	"cloud_disk/go-zreo_cloud-disk/core/helper"
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"cloud_disk/go-zreo_cloud-disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreatereq, userIdentity string) (resp *types.UserFolderCreateres, err error) {
	// 判断新名称在当前目录是否存在
	// 子查询获取parentid,再判断
	cnt, err := l.svcCtx.Xorm.Where("name =? and parent_id= ? ", req.Name, req.ParentId).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("file has been existing")
	}
	// 创建文件夹
	data := &models.UserRepository{
		Identity:     helper.GenerateUUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Ext:          "",
		Name:         req.Name,
	}
	_, err = l.svcCtx.Xorm.Insert(data)
	if err != nil {
		return nil, err
	}
	return
}
