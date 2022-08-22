package repository

import (
	"git.imooc.com/keegan/product/domain/model"
	"github.com/jinzhu/gorm"
	)

type IProductRepository interface {
	InitTable() error
	FindProductByID(int64) (*model.Product,error)
	CreateProduct(*model.Product)(int64,error)
	DeleteProductByID(int64) error
	UpdateProduct(*model.Product)error
	FindAll()([]model.Product,error)
}

// 创建productRepository
func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDB:db}
}

type ProductRepository struct {
	mysqlDB *gorm.DB
}

// 初始化4张表
func (u *ProductRepository)InitTable()error{
	return u.mysqlDB.CreateTable(&model.Product{},&model.ProductSeo{},&model.ProductImage{},&model.ProductSize{}).Error
}

// 创建Product信息
func (u *ProductRepository) CreateProduct(product *model.Product)(int64,error){
	return product.ID,u.mysqlDB.Create(product).Error
}

// 根据ID删除Product信息
func (u *ProductRepository)DeleteProductByID(productID int64) error{
	// 开启事务
	tx := u.mysqlDB.Begin()
	defer func() {
		if r:=recover();r!=nil{
			tx.Rollback()
		}
	}()

	if tx.Error != nil{
		return tx.Error
	}

	// 删除
	if err := tx.Unscoped().Where("id=?",productID).Delete(&model.Product{}).Error;err!=nil{
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("images_product_id=?",productID).Delete(&model.ProductImage{}).Error;err!=nil{
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("size_product_id=?",productID).Delete(&model.ProductSize{}).Error;err!=nil{
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("seo_product_id=?",productID).Delete(&model.ProductSeo{}).Error;err!=nil{
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 更新Product信息
func (u *ProductRepository) UpdateProduct(product *model.Product) error{
	return u.mysqlDB.Model(product).Update(product).Error
}

// 根据ID查找Product信息
func (u *ProductRepository)FindProductByID(productID int64)(product *model.Product,err error){
	product = &model.Product{}
	return product ,u.mysqlDB.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").First(product,productID).Error
}

// 获取结果集
func (u *ProductRepository)FindAll()(productAll []model.Product,err error){
	return productAll,u.mysqlDB.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").Find(&productAll).Error
}