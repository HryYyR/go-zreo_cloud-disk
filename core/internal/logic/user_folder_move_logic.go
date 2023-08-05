package logic

import (
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"cloud_disk/go-zreo_cloud-disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderMoveLogic {
	return &UserFolderMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderMoveLogic) UserFolderMove(req *types.UserFolderMovereq, userIdentity string) (resp *types.UserFolderMoveres, err error) {
	parentDate := new(models.UserRepository)
	has, err := l.svcCtx.Xorm.Where("identity=? and user_identity=?", req.ParentIdentity, userIdentity).Get(parentDate)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("folder does not exist")
	}

	// update
	_, err = l.svcCtx.Xorm.Where("identity=?", req.Identity).Update(models.UserRepository{ParentId: parentDate.Id})
	if err != nil {
		return nil, err
	}
	return
}
