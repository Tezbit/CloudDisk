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

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   repository.IUserRepoRepository
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repository.NewUserRepoRepository(svcCtx.Engine),
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReponse, err error) {
	//parentID
	parentData := new(models.UserRepository)
	has, err := l.repo.FindUserRepoByIdentityAndUserIdentity(req.ParentIdentity, userIdentity, parentData)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("文件夹不存在")
	}
	// 更新记录的 ParentID
	err = l.repo.UpdateUserRepoByIdentity(req.Idnetity, &models.UserRepository{
		ParentId: int64(parentData.Id),
	})
	return
}
