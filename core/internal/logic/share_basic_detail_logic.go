package logic

import (
	"cloud_disk/core/domain/repository"
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	sbRepo repository.IShareBasicRepository
	asRepo repository.IAssociateRepository
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		sbRepo: repository.NewShareBasicRepository(svcCtx.Engine),
		asRepo: repository.NewAssociateRepository(svcCtx.Engine),
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailResponse, err error) {
	// 对分享记录的点击次数进行 + 1
	err = l.sbRepo.UpdateByShareBasicIdentity(req.Identity)
	if err != nil {
		return
	}
	// 获取资源的详细信息
	data, err := l.asRepo.GetShareBasicDetails(req.Identity)
	resp = &types.ShareBasicDetailResponse{
		RepositoryIdentity: data.RepositoryIdentity,
		Name:               data.Name,
		Ext:                data.Ext,
		Size:               data.Size,
		Path:               data.Path,
	}
	return
}
