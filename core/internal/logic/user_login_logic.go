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

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.Loginreq) (resp *types.Loginres, err error) {
	user := new(models.UserBasic)
	//查询用户是否存在
	ext, err := l.svcCtx.Xorm.Where("name=? and password =?", req.Name, helper.Md5(req.PassWord)).Get(user)
	// ext, err := models.Xorm.Where("name=? and password =?", req.Name, req.PassWord).Get(user)
	if err != nil {
		return nil, err
	}
	if !ext {
		return nil, errors.New("用户名或密码错误!")
	}
	// 生成token
	token, err := helper.GenerateToken(uint64(user.Id), user.Identity, user.Name)
	if err != nil {
		return nil, errors.New("generate token error: " + err.Error())
	}
	return &types.Loginres{Token: token}, nil
}
