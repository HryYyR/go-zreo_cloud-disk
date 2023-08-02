package logic

import (
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"cloud_disk/go-zreo_cloud-disk/core/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderDeleteLogic {
	return &UserFolderDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderDeleteLogic) UserFolderDelete(req *types.UserFolderDeletereq, userIdentity string) (resp *types.UserFolderDeleteres, err error) {
	data := &models.UserRepository{
		Identity:     req.Identity,
		UserIdentity: userIdentity,
	}
	_, err = l.svcCtx.Xorm.Delete(data)
	if err != nil {
		return nil, err
	}

	return
}
