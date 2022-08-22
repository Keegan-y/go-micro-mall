package service

import (
	"errors"
	"fmt"
	"git.imooc.com/keegan/user/domain/model"
	"git.imooc.com/keegan/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(*model.User) (int64,error)
	DeleteUser(int64) error
	UpdateUser(user *model.User,isChangePwd bool)(err error)
	FindUserByName(string) (*model.User,error)
	CheckPwd(userName string,pwd string) (isOK bool,err error)
}

// 创建实例
func NewUserDataService(userRepository repository.IUserRepository)IUserDataService  {
	return &UserDataService{UserRepository:userRepository}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}
// 加密用户密码
func GeneratePassword(userPassword string) ([]byte,error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}
// 验证用户密码
func ValidatePassword(userPassword string,hashed string) (isOK bool,err error){
	if err = bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(userPassword));err !=nil{
		return false,errors.New("密码对比错误")
	}
	return true,nil
}
// AddUser 新增用户
func (u *UserDataService) AddUser(user *model.User) (userID int64,err error){
	pwdByte,err := GeneratePassword(user.HashPassword)
	if err !=nil{
		return user.ID,err
	}
	user.HashPassword = string(pwdByte)
	// 业务逻辑，比如rabbitmq
	return u.UserRepository.CreateUser(user)
}
// 删除用户
func (u *UserDataService)DeleteUser(userID int64) error{
	return u.UserRepository.DeleteUserByID(userID)
}
// 更新用户
func (u *UserDataService)UpdateUser(user *model.User,isChangePwd bool)(err error){
	// 判断是否更新了代码
	if isChangePwd {
		pwdByte,err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	// 业务逻辑，比如log
	return u.UserRepository.UpdateUser(user)
}
// FindUserByName 根据用户名称查用户
func (u *UserDataService)FindUserByName(userName string)(user *model.User,err error){
	fmt.Printf("user数据类型%T\n",user)
	return u.UserRepository.FindUserByName(userName)
}
// 对比账号密码是否一致
func (u *UserDataService)CheckPwd(userName string,pwd string) (isOK bool,err error){
	user,err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false,err
	}
	return ValidatePassword(pwd,user.HashPassword)
}
