package logic

import (
	"cloud_disk/core/define"
	"cloud_disk/core/domain/models"
	"cloud_disk/core/domain/repository"
	"cloud_disk/core/helper"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   repository.IRepoPoolRepository
}

func NewFileUploadChunkCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkCompleteLogic {
	return &FileUploadChunkCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repository.NewRepoPoolRepository(svcCtx.Engine),
	}
}

func (l *FileUploadChunkCompleteLogic) FileUploadChunkComplete(req *types.FileUploadChunkCompleteRequest) (resp *types.FileUploadChunkCompleteResponse, err error) {
	co := make([]cos.Object, 0)
	for _, v := range req.CosObjects {
		co = append(co, cos.Object{
			ETag:       v.Etag,
			PartNumber: v.PartNumber,
		})
	}
	err = helper.CosPartUploadComplete(l.svcCtx, req.Key, req.UploadId, co)

	// 数据入库
	rp := &models.RepositoryPool{
		Identity: helper.GetUUid(),
		Hash:     req.Md5Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     define.Bucket + "/" + req.Key,
	}
	err = l.repo.CreateRepoPool(rp)
	if err != nil {
		return nil, err
	}
	return
}
