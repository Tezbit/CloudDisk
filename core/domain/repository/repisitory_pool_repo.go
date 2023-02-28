package repository

import (
	"cloud_disk/core/domain/models"
	"xorm.io/xorm"
)

type IRepoPoolRepository interface {
	FindRepoPoolByHash(hash string) (*models.RepositoryPool, error)
	CreateRepoPool(*models.RepositoryPool) error
	GetRepoPoolByIdentity(string, *models.RepositoryPool) (bool, error)
}

func NewRepoPoolRepository(engine *xorm.Engine) IRepoPoolRepository {
	return &RepoPoolRepository{
		db: engine,
	}
}

type RepoPoolRepository struct {
	db *xorm.Engine
}

func (r *RepoPoolRepository) GetRepoPoolByIdentity(identity string, rp *models.RepositoryPool) (bool, error) {
	return r.db.Where("identity = ?", identity).Get(rp)
}

func (r *RepoPoolRepository) CreateRepoPool(rp *models.RepositoryPool) error {
	_, err := r.db.Insert(rp)
	return err
}

func (r *RepoPoolRepository) FindRepoPoolByHash(hash string) (*models.RepositoryPool, error) {
	rp := &models.RepositoryPool{}
	_, err := r.db.Where("hash=?", hash).Get(rp)
	return rp, err
}
