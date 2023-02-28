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

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	rpRepo repository.IRepoPoolRepository
	urRepo repository.IUserRepoRepository
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		rpRepo: repository.NewRepoPoolRepository(svcCtx.Engine),
		urRepo: repository.NewUserRepoRepository(svcCtx.Engine),
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveResponse, err error) {
	// 获取资源详情
	rp := new(models.RepositoryPool)
	has, err := l.rpRepo.GetRepoPoolByIdentity(req.RepositoryIdentity, rp)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("资源不存在")
	}
	// user_repository 资源保存
	ur := &models.UserRepository{
		Identity:           helper.GetUUid(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}
	err = l.urRepo.CreateUserRepo(ur)
	resp = new(types.ShareBasicSaveResponse)
	resp.Identity = ur.Identity
	return
}
