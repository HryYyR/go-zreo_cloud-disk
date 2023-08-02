package logic

import (
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"cloud_disk/go-zreo_cloud-disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdatereq, userIdentity string) (resp *types.UserFileNameUpdateres, err error) {
	// 判断新名称在当前目录是否存在
	// 子查询获取parentid,再判断
	cnt, err := l.svcCtx.Xorm.
		Where("name =? and parent_id= ( SELECT parent_id FROM user_repository ur WHERE ur.repository_identity =? ) ", req.Name, req.Identity).
		Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("file has been existing")
	}
	udata := &models.UserRepository{Name: req.Name}
	rdata := &models.RepositoryPool{Name: req.Name}
	session := l.svcCtx.Xorm.NewSession()
	err = session.Begin()
	if err != nil {
		return nil, err
	}
	_, err = session.Where("repository_identity =? and user_identity =? ", req.Identity, userIdentity).Update(udata)
	if err != nil {
		session.Rollback()
		return
	}
	_, err = session.Where("identity =?", req.Identity).Update(rdata)
	if err != nil {
		session.Rollback()
		return
	}
	err = session.Commit()
	if err != nil {
		return nil, err
	}
	return
}
