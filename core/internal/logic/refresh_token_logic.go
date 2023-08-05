package logic

import (
	"cloud_disk/go-zreo_cloud-disk/core/helper"
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenreq, authorization string) (resp *types.RefreshTokenres, err error) {
	uc, err := helper.DecryptToken(authorization)
	if err != nil {
		return nil, err
	}
	token, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, time.Hour*12)
	if err != nil {
		return nil, err
	}
	refreshtoken, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, time.Hour*84)
	if err != nil {
		return nil, err
	}
	resp = new(types.RefreshTokenres)
	resp.Token = token
	resp.RefreshToken = refreshtoken
	return
}
