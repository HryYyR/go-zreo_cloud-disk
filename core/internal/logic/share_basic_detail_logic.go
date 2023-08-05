package logic

import (
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailreq) (resp *types.ShareBasicDetailres, err error) {
	session := l.svcCtx.Xorm.NewSession()
	session.Begin()
	//更新share的点击次数
	_, err = session.Exec("UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ?", req.Identity)
	if err != nil {
		session.Rollback()
		return nil, err
	}
	// 获取share详情
	resp = new(types.ShareBasicDetailres)
	_, err = session.Table("share_basic").
		Select("share_basic.repository_identity , repository_pool.name,repository_pool.ext,repository_pool.size,repository_pool.path").
		Join("LEFT", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
		Where("share_basic.identity = ?", req.Identity).Get(resp)
	if err != nil {
		session.Rollback()
		return nil, err
	}
	session.Commit()
	return resp, nil
}
