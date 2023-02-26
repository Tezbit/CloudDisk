package logic

import (
	repo "cloud_disk/core/domain/repository"
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	userBasicRepo repo.IUserBasicRepository
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		userBasicRepo: repo.NewUserBasicRepository(svcCtx.Engine),
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	user, err := l.userBasicRepo.FindUserByIdentity(req.Identity)
	if err != nil {
		return nil, err
	}
	resp = &types.UserDetailResponse{
		Name:  user.Name,
		Email: user.Email,
	}
	return
}
