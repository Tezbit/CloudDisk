package logic

import (
	"cloud_disk/core/domain/models"
	repo "cloud_disk/core/domain/repository"
	"cloud_disk/core/helper"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	userBasicRepo repo.IUserBasicRepository
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		userBasicRepo: repo.NewUserBasicRepository(svcCtx.Engine),
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("该邮箱验证码为空")
	}
	if code != req.Code {
		return nil, errors.New("验证码不一致")
	}
	count, err := l.userBasicRepo.GetNameCount(req.Name)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("用户已存在")
	}
	//数据入库
	gPwd, err := helper.Generate(req.Password)
	if err != nil {
		return nil, err
	}
	user := &models.UserBasic{
		Identity: helper.GetUUid(),
		Name:     req.Name,
		Password: gPwd,
		Email:    req.Email,
	}
	err = l.userBasicRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	//注册成功删除验证码
	l.svcCtx.RDB.Del(l.ctx, req.Email)
	return
}
