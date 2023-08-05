package logic

import (
	"cloud_disk/go-zreo_cloud-disk/core/helper"
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"cloud_disk/go-zreo_cloud-disk/core/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareCreateLogic {
	return &ShareCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareCreateLogic) ShareCreate(req *types.ShareCreatereq, userIdentity string) (resp *types.ShareCreateres, err error) {
	uuid := helper.GenerateUUID()
	data := models.ShareBasic{
		Identity:           uuid,
		UserIdentity:       userIdentity,
		RepositoryIdentity: req.RepositoryIdentity,
		ExpiredTime:        req.ExpiredTime,
		ClickNum:           0,
	}
	_, err = l.svcCtx.Xorm.Insert(data)
	if err != nil {
		return nil, err
	}
	resp = &types.ShareCreateres{ShareIdentity: uuid}
	return resp, nil
}
