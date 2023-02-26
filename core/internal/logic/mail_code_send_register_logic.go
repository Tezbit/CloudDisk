package logic

import (
	"cloud_disk/core/define"
	repo "cloud_disk/core/domain/repository"
	"cloud_disk/core/helper"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	userBasicRepo repo.IUserBasicRepository
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		userBasicRepo: repo.NewUserBasicRepository(svcCtx.Engine),
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeRequest) (resp *types.MailCodeResponse, err error) {
	//该邮箱未注册
	count, err := l.userBasicRepo.GetEmailCount(req.Email)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("该邮箱已被注册")
	}
	//验证码
	_, err = l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err == nil {
		//验证码有效期内
		return nil, errors.New("验证码尚在有效期")
	}
	code := helper.Captcha()
	err = l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CondeExpire)).Err()
	if err != nil {
		return nil, err
	}
	//todo 异步？
	sender := l.svcCtx.Config.Sender
	err = helper.MailSendCode(sender.Email, req.Email, code, sender.AuthPwd)
	if err != nil {
		return nil, err
	}
	resp = new(types.MailCodeResponse)
	resp.Code = code
	return
}
