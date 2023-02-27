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

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   repository.IUserRepoRepository
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repository.NewUserRepoRepository(svcCtx.Engine),
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateResponse, err error) {
	// 判断当前名称在该层级下是否存在
	cnt, err := l.repo.GetUserRepoCountByNameAndUserIdentityAndParentId(req.Name, userIdentity, int(req.ParentId))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("该名称已存在")
	}
	//创建文件夹
	ur := &models.UserRepository{
		Identity:     helper.GetUUid(),
		UserIdentity: userIdentity,
		ParentId:     int64(req.ParentId),
		Name:         req.Name,
	}

	err = l.repo.CreateUserRepo(ur)
	if err != nil {
		return nil, err
	}
	resp = new(types.UserFolderCreateResponse)
	resp.Identity = ur.Identity
	return

}
