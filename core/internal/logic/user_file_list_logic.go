package logic

import (
	"cloud_disk/core/define"
	"cloud_disk/core/domain/models"
	"cloud_disk/core/domain/repository"
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	repo     repository.IAssociateRepository
	userRepo repository.IUserRepoRepository
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		repo:     repository.NewAssociateRepository(svcCtx.Engine),
		userRepo: repository.NewUserRepoRepository(svcCtx.Engine),
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListResponse, err error) {
	uf := make([]*models.UserFile, 0)
	resp = new(types.UserFileListResponse)
	// 分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	ur, err := l.userRepo.FindUserRepoIdByIdentity(req.Identity)
	if err != nil {
		return nil, err
	}
	err = l.repo.GetUserFileList(ur.Id, userIdentity, size, offset, uf)
	if err != nil {
		return nil, err
	}
	//todo: 查询某一层级用户文件总数 剔除文件夹？
	cnt, err := l.userRepo.GetUserRepoFileCountByParentIdAndUserIdentity(ur.Id, userIdentity)
	if err != nil {
		return nil, err
	}
	resp.List = make([]*types.UserFile, 0)
	for _, file := range uf {
		resp.List = append(resp.List, &types.UserFile{
			Id:                 file.Id,
			Identity:           file.Identity,
			RepositoryIdentity: file.RepositoryIdentity,
			Name:               file.Name,
			Ext:                file.Ext,
			Path:               file.Path,
			Size:               file.Size,
		})
	}
	resp.Count = cnt
	return
}
