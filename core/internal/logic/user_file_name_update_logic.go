package logic

import (
	"cloud_disk/core/domain/models"
	"cloud_disk/core/domain/repository"
	"context"
	"errors"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   repository.IUserRepoRepository
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repository.NewUserRepoRepository(svcCtx.Engine),
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReponse, err error) {
	// 判断当前名称在该层级下是否存在
	cnt, err := l.repo.GetUserRepoCountByNameAndUserIdentityAndIdentity(req.Name, userIdentity, req.Identity)
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("该名称已存在")
	}
	// 文件名称修改
	data := &models.UserRepository{Name: req.Name}
	err = l.repo.UpdateUserRepo(req.Identity, userIdentity, data)
	if err != nil {
		return
	}
	return
}
