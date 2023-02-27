package logic

import (
	"cloud_disk/core/domain/models"
	"cloud_disk/core/domain/repository"
	"cloud_disk/core/helper"
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx                context.Context
	svcCtx             *svc.ServiceContext
	userRepositoryRepo repository.IUserRepoRepository
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger:             logx.WithContext(ctx),
		ctx:                ctx,
		svcCtx:             svcCtx,
		userRepositoryRepo: repository.NewUserRepoRepository(svcCtx.Engine),
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, userIdentity string) (resp *types.UserRepositorySaveResponse, err error) {
	ur := &models.UserRepository{
		Identity:           helper.GetUUid(),
		UserIdentity:       userIdentity,
		ParentId:           int64(req.ParentId),
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	// 通过接口调用(upload/在前)来check重复文件问题
	err = l.userRepositoryRepo.CreateUserRepo(ur)
	if err != nil {
		return nil, err
	}

	return
}
