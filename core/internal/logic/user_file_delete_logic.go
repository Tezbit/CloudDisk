package logic

import (
	"cloud_disk/core/domain/repository"
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	rpRepo repository.IRepoPoolRepository
	urRepo repository.IUserRepoRepository
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		rpRepo: repository.NewRepoPoolRepository(svcCtx.Engine),
		urRepo: repository.NewUserRepoRepository(svcCtx.Engine),
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest, userIdentity string) (resp *types.UserFileDeleteReponse, err error) {
	// 删除用户数据
	// 软删除，不删除repository_pool库的条目
	err = l.urRepo.DeleteByUserIdentityAndIdentity(userIdentity, req.Identity)
	return
}
