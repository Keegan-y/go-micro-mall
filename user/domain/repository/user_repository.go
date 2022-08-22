package repository

import (
	"git.imooc.com/keegan/user/domain/model"
	"github.com/jinzhu/gorm"
	)

type IUserRepository interface {
	//初始化数据表
	InitTable() error
	//根据用户名称查找用户信息
	FindUserByName(string) (*model.User,error)
	//根据用户ID查找用户信息
	FindUserByID(int64) (*model.User,error)
	//创建用户
	CreateUser(*model.User) (int64,error)
	//根据用户ID删除用户
	DeleteUserByID(int64) error
	// UpdateUser 更新用户信息
	UpdateUser(*model.User) error
	// FindAll 查找所有用
	FindAll() ([]model.User,error)
}

// 创建UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository  {
	return &UserRepository{mysqlDB:db}
}

type UserRepository struct {
	mysqlDB *gorm.DB
}

// 初始化表
func (u *UserRepository) InitTable() error {
	return u.mysqlDB.CreateTable(&model.User{}).Error
}

// 根据用户名查找用户信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user,u.mysqlDB.Where("user_name=?",name).Find(user).Error
}

// 根据用户ID查找用户信息
func (u *UserRepository) FindUserByID(userID int64) (user *model.User,err error) {
	user = &model.User{}
	return user,u.mysqlDB.First(user,userID).Error
}

// 创建用户
func (u *UserRepository) CreateUser(user *model.User) (userID int64,err error) {
	return user.ID,u.mysqlDB.Create(user).Error
}

// 根据用户ID删除用户
func (u *UserRepository) DeleteUserByID(userID int64) error {
	return u.mysqlDB.Where("id=?",userID).Error
}


// 更新用户信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDB.Model(user).Update(&user).Error
}

// 查找所有用户
func (u *UserRepository) FindAll() (userAll []model.User,err error) {
	return userAll, u.mysqlDB.Find(&userAll).Error
}
