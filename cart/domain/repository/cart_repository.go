package repository

import (
	"errors"
	"git.imooc.com/keegan/cart/domain/model"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart,error)
	CreateCart(*model.Cart)(int64,error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64)([]model.Cart,error)

	CleanCart(int64)error
	IncrNum(int64,int64)error
	DecrNum(int64,int64)error
}

// 创建cartRepository
func NewCartRepository(db *gorm.DB)ICartRepository{
	return &CartRepository{mysqlDB:db}
}

type CartRepository struct {
	mysqlDB *gorm.DB
}

// 初始化表
func (u *CartRepository)InitTable()error{
	return u.mysqlDB.CreateTable(&model.Cart{}).Error
}

// 创建Cart信息
func(u *CartRepository)CreateCart(cart *model.Cart)(int64,error){
	db := u.mysqlDB.FirstOrCreate(cart,model.Cart{ProductID:cart.ProductID,SizeID:cart.SizeID,UserID:cart.UserID})
	if db.Error != nil{
		return 0,db.Error
	}
	if db.RowsAffected == 0{
		return 0,errors.New("购物车插入失败")
	}
	return cart.ID,nil
}

// 添加商品数量
func(u *CartRepository)IncrNum(cartID int64,num int64)error{
	cart := &model.Cart{ID:cartID}
	return u.mysqlDB.Model(cart).UpdateColumn("num",gorm.Expr("num+?",num)).Error
}

// 根据ID删除Cart信息
func(u *CartRepository)DeleteCartByID(cartID int64)error{
	return u.mysqlDB.Where("id=?",cartID).Delete(&model.Cart{}).Error
}

// 根据用户ID来清空购物车
func(u *CartRepository)CleanCart(userID int64)error{
	return u.mysqlDB.Where("user_id=?",userID).Delete(&model.Cart{}).Error
}

// 购物车减少商品
func(u *CartRepository)DecrNum(cartID int64,num int64)error{
	cart := &model.Cart{ID:cartID}
	db := u.mysqlDB.Model(cart).Where("num >= ?",num).UpdateColumn("num",gorm.Expr("num - ?",num))
	if db.Error != nil{
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}

// 更新Cart信息
func(u *CartRepository)UpdateCart(cart *model.Cart)error{
	return u.mysqlDB.Model(cart).Update(cart).Error
}

// 根据ID查找Cart信息
func(u *CartRepository)FindCartByID(cartID int64)(cart *model.Cart,err error){
	cart = &model.Cart{}
	return cart,u.mysqlDB.First(cart,cartID).Error
}

// 获取结果集
func (u *CartRepository)FindAll(userID int64)(cartAll []model.Cart,err error){
	return cartAll,u.mysqlDB.Where("user_id=?",userID).Find(&cartAll).Error
}