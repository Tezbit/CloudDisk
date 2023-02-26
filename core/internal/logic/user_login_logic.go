package logic

import (
	repo "cloud_disk/core/domain/repository"
	"cloud_disk/core/helper"
	"context"
	"errors"
	"fmt"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	userBasicRepo repo.IUserBasicRepository
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		userBasicRepo: repo.NewUserBasicRepository(svcCtx.Engine),
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	//查询用户
	fmt.Println(req.Name)
	user, err := l.userBasicRepo.FindUserByName(req.Name)
	if err != nil {
		return nil, err
	}
	check, err := helper.Validate(req.Password, user.Password)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, errors.New("密码错误")
	}
	//生成token
	token, err := helper.GenerateToken(uint64(user.Id), user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResponse)
	resp.Tocken = token
	return
}
