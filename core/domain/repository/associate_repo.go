package repository

import (
	"cloud_disk/core/define"
	"cloud_disk/core/domain/models"
	"time"
	"xorm.io/xorm"
)

type IAssociateRepository interface {
	GetUserFileList(int, string, int, int, []*models.UserFile) error
	GetShareBasicDetails(string) (*models.ShareBasicDetail, error)
}

func NewAssociateRepository(engine *xorm.Engine) IAssociateRepository {
	return &AssociateRepository{
		db: engine,
	}
}

type AssociateRepository struct {
	db *xorm.Engine
}

func (a *AssociateRepository) GetShareBasicDetails(sbIdentity string) (*models.ShareBasicDetail, error) {
	sbd := &models.ShareBasicDetail{}
	_, err := a.db.Table("share_basic").
		Select("share_basic.repository_identity, user_repository.name, repository_pool.ext, repository_pool.size, repository_pool.path").
		Join("LEFT", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
		Join("LEFT", "user_repository", "user_repository.identity = share_basic.user_repository_identity").
		Where("share_basic.identity = ?", sbIdentity).Get(sbd)
	if err != nil {
		return nil, err
	}
	return sbd, nil
}

func (a *AssociateRepository) GetUserFileList(parentId int, userIdentity string, size int, offset int, uf []*models.UserFile) error {
	return a.db.Table("user_repository").Where("parent_id = ? AND user_identity = ? ", parentId, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.ext,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Limit(size, offset).Find(&uf)
}
