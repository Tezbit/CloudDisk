package repository

import (
	"cloud_disk/core/domain/models"
	"xorm.io/xorm"
)

type IUserBasicRepository interface {
	FindUserByName(string) (*models.UserBasic, error)
	FindUserByIdentity(string) (*models.UserBasic, error)
	GetEmailCount(string) (int64, error)
	GetNameCount(string) (int64, error)
	CreateUser(*models.UserBasic) error
}

func NewUserBasicRepository(engine *xorm.Engine) IUserBasicRepository {
	return &UserBasicRepository{
		db: engine,
	}
}

type UserBasicRepository struct {
	db *xorm.Engine
}

func (u *UserBasicRepository) CreateUser(user *models.UserBasic) error {
	_, err := u.db.Insert(user)
	return err
}

func (u *UserBasicRepository) GetNameCount(name string) (int64, error) {
	return u.db.Where("name=?", name).Count(new(models.UserBasic))
}

func (u *UserBasicRepository) GetEmailCount(email string) (int64, error) {
	return u.db.Where("email=?", email).Count(new(models.UserBasic))
}

func (u *UserBasicRepository) FindUserByIdentity(identity string) (*models.UserBasic, error) {
	user := &models.UserBasic{}
	_, err := u.db.Where("identity=?", identity).Get(user)
	//if !ok {
	//	return nil, errors.New("查询的对象不存在")
	//}
	return user, err
}

func (u *UserBasicRepository) FindUserByName(name string) (*models.UserBasic, error) {
	user := &models.UserBasic{}
	_, err := u.db.Where("name=?", name).Get(user)
	//if !ok {
	//	return nil, errors.New("查询的对象不存在")
	//}
	return user, err
}
