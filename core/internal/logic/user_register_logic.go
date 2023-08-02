package logic

import (
	"cloud_disk/go-zreo_cloud-disk/core/helper"
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"cloud_disk/go-zreo_cloud-disk/core/models"
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.Userregisterreq) (resp *types.Userregisterres, err error) {
	code := l.svcCtx.Redis.Get(req.Email).Val()
	fmt.Println(code)
	// 判断验证码是否一致
	if code != req.Code {
		return nil, errors.New("Invalid code")
	}
	// 判断用户名或者邮箱是否已被占用
	count, err := l.svcCtx.Xorm.Where("name=? or email=?", req.Username, req.Email).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if count != 0 {
		return nil, errors.New("the username or email has already been used. Please try again")
	}

	user := &models.UserBasic{
		Identity: helper.GenerateUUID(),
		Name:     req.Username,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	fmt.Printf("%v", user)

	_, err = l.svcCtx.Xorm.InsertOne(user)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
