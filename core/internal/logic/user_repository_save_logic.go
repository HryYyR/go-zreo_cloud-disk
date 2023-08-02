package logic

import (
	"cloud_disk/go-zreo_cloud-disk/core/helper"
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"cloud_disk/go-zreo_cloud-disk/core/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySavereq, userIdentity string) (resp *types.UserRepositorySaveres, err error) {
	ur := &models.UserRepository{
		Identity:           helper.GenerateUUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	_, err = l.svcCtx.Xorm.Insert(ur)
	if err != nil {
		return
	}
	return &types.UserRepositorySaveres{
		Identity: ur.Identity,
	}, nil
}
