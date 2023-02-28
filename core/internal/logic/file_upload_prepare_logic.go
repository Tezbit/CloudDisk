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

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   repository.IRepoPoolRepository
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repository.NewRepoPoolRepository(svcCtx.Engine),
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareResponse, err error) {
	rp := new(models.RepositoryPool)
	has, err := l.repo.FindRepoPoolByHash(req.Md5Hash, rp)
	if err != nil {
		return
	}
	resp = new(types.FileUploadPrepareResponse)
	if has {
		//已经有了就直接返回
		resp.Identity = rp.Identity
	} else {
		//准备分片上传
		key, uploadId, err := helper.CosInitPart(l.svcCtx, req.Ext)
		if err != nil {
			return nil, err
		}
		resp.Key = key
		resp.UploadId = uploadId
	}

	return
}
