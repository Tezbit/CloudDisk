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

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	rpRepo repository.IRepoPoolRepository
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		rpRepo: repository.NewRepoPoolRepository(svcCtx.Engine),
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadResponse, err error) {
	rp := &models.RepositoryPool{
		Identity: helper.GetUUid(),
		Hash:     req.Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     req.Path,
	}
	err = l.rpRepo.CreateRepoPool(rp)
	if err != nil {
		return nil, err
	}
	resp = &types.FileUploadResponse{
		Identity: rp.Identity,
		Ext:      rp.Ext,
		Name:     rp.Name,
	}
	return
}
