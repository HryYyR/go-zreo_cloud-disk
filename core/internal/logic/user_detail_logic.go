package logic

import (
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"cloud_disk/go-zreo_cloud-disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.Userdetailreq) (resp *types.Userdetailres, err error) {
	resp = &types.Userdetailres{}
	ub := new(models.UserBasic)
	has, err := l.svcCtx.Xorm.Where("identity =?", req.Identity).Get(ub)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	resp.Name = ub.Name
	resp.Email = ub.Email
	// b, err := json.Marshal(data)
	// dst := new(bytes.Buffer)
	// err = json.Indent(dst, b, "", " ")
	// if err != nil {
	// 	log.Println("json indent error:", err)
	// }

	return resp, nil
}
