package logic

import (
	"cloud_disk/go-zreo_cloud-disk/core/helper"
	"cloud_disk/go-zreo_cloud-disk/core/internal/svc"
	"cloud_disk/go-zreo_cloud-disk/core/internal/types"
	"cloud_disk/go-zreo_cloud-disk/core/models"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendreq) (resp *types.MailCodeSendres, err error) {
	cnt, err := l.svcCtx.Xorm.Where("email=?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("Mail have already been registered")
	}

	mail := l.svcCtx.Redis.Get(req.Email).Val()
	if len(mail) != 0 {
		return nil, errors.New("you have already send the code within 5 minutes, please try again later")
	}
	code := helper.RandCode()
	// var wg sync.WaitGroup
	// wg.Add(2)
	go func() {
		// defer wg.Done()
		l.svcCtx.Redis.Set(req.Email, code, 5*time.Minute) //验证码存redis
	}()
	go func(vcode string) {
		// defer wg.Done()
		err = helper.MailSendCode(req.Email, vcode) //发送验证码
		if err != nil {
			fmt.Println("code send error", err)
		}
	}(code)
	// wg.Wait()
	return
}
