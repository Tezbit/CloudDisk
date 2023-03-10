package repository

import (
	"cloud_disk/core/domain/models"
	"xorm.io/xorm"
)

type IUserRepoRepository interface {
	CreateUserRepo(*models.UserRepository) error
	FindUserRepoIdByIdentity(string) (*models.UserRepository, error)
	GetUserRepoFileCountByParentIdAndUserIdentity(int, string) (int64, error)
	GetUserRepoCountByNameAndUserIdentityAndIdentity(string, string, string) (int64, error)
	UpdateUserRepoByIdentityAndUserIdentity(string, string, *models.UserRepository) error
	GetUserRepoCountByNameAndUserIdentityAndParentId(string, string, int) (int64, error)
	FindUserRepoByIdentityAndUserIdentity(string, string, *models.UserRepository) (bool, error)
	DeleteByUserIdentityAndIdentity(string, string) error
	UpdateUserRepoByIdentity(string, *models.UserRepository) error
	GetUserRepoByIdentity(string, *models.UserRepository) (bool, error)
}

func NewUserRepoRepository(engine *xorm.Engine) IUserRepoRepository {
	return &UserRepoRepository{
		db: engine,
	}
}

type UserRepoRepository struct {
	db *xorm.Engine
}

func (u *UserRepoRepository) GetUserRepoByIdentity(identity string, data *models.UserRepository) (bool, error) {
	return u.db.Where("identity = ?", identity).Get(data)
}

func (u *UserRepoRepository) UpdateUserRepoByIdentity(identity string, data *models.UserRepository) error {
	_, err := u.db.Where("identity = ?", identity).Update(data)
	return err
}

func (u *UserRepoRepository) FindUserRepoByIdentityAndUserIdentity(identity string, userIdentity string, data *models.UserRepository) (bool, error) {
	return u.db.Where("identity = ? AND user_identity = ?", identity, userIdentity).Get(data)
}

func (u *UserRepoRepository) DeleteByUserIdentityAndIdentity(userIdentity string, identity string) error {
	_, err := u.db.Where("user_identity = ? AND identity = ?", userIdentity, identity).Delete(new(models.UserRepository))
	return err
}

func (u *UserRepoRepository) GetUserRepoCountByNameAndUserIdentityAndParentId(name string, userIdentity string, parentId int) (int64, error) {
	return u.db.Where("name = ? AND user_identity = ? AND parent_id = ?", name, userIdentity, parentId).Count(new(models.UserRepository))
}

func (u *UserRepoRepository) GetUserRepoCountByNameAndUserIdentityAndIdentity(name string, userIdentity string, identity string) (int64, error) {
	// ????????????????????????????????????????????????????????????
	return u.db.Where("name = ? AND user_identity = ? AND parent_id = (SELECT parent_id FROM user_repository ur WHERE ur.identity = ?)", name, userIdentity, identity).Count(new(models.UserRepository))
}

func (u *UserRepoRepository) UpdateUserRepoByIdentityAndUserIdentity(identity string, userIdentity string, data *models.UserRepository) error {
	_, err := u.db.Where("identity = ? AND user_identity = ? ", identity, userIdentity).Update(data)
	return err
}

func (u *UserRepoRepository) GetUserRepoFileCountByParentIdAndUserIdentity(id int, userIdentity string) (int64, error) {
	return u.db.Where("parent_id = ? AND user_identity = ? ", id, userIdentity).Count(new(models.UserRepository))
}

func (u *UserRepoRepository) FindUserRepoIdByIdentity(identity string) (*models.UserRepository, error) {
	ur := &models.UserRepository{}
	_, err := u.db.Table("user_repository").Select("id").
		Where("identity = ?", identity).Get(ur)
	if err != nil {
		return nil, err
	}
	return ur, nil
}

func (u *UserRepoRepository) CreateUserRepo(ur *models.UserRepository) error {
	_, err := u.db.Insert(ur)
	return err
}
