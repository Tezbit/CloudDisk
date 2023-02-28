package logic

import (
	"cloud_disk/core/domain/models"
	"cloud_disk/core/domain/repository"
	"cloud_disk/core/helper"
	"context"
	"errors"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	urRepo repository.IUserRepoRepository
	sbRepo repository.IShareBasicRepository
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		urRepo: repository.NewUserRepoRepository(svcCtx.Engine),
		sbRepo: repository.NewShareBasicRepository(svcCtx.Engine),
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateResponse, err error) {
	uuid := helper.GetUUid()
	ur := new(models.UserRepository)
	has, err := l.urRepo.GetUserRepoByIdentity(req.UserRepositoryIdentity, ur)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user repository not found")
	}

	data := &models.ShareBasic{
		Identity:               uuid,
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     ur.RepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}
	err = l.sbRepo.Create(data)
	if err != nil {
		return
	}
	resp = &types.ShareBasicCreateResponse{
		Identity: uuid,
	}

	return
}
