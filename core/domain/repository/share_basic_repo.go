package repository

import (
	"cloud_disk/core/domain/models"
	"xorm.io/xorm"
)

type IShareBasicRepository interface {
	Create(*models.ShareBasic) error
	UpdateByShareBasicIdentity(string) error
}

func NewShareBasicRepository(engine *xorm.Engine) IShareBasicRepository {
	return &ShareBasicRepository{
		db: engine,
	}
}

type ShareBasicRepository struct {
	db *xorm.Engine
}

func (s *ShareBasicRepository) UpdateByShareBasicIdentity(identity string) error {
	_, err := s.db.Exec("UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ?", identity)
	return err
}

func (s *ShareBasicRepository) Create(data *models.ShareBasic) error {
	_, err := s.db.Insert(data)
	return err
}
