package service

import (
	"git.imooc.com/keegan/product/domain/model"
	"git.imooc.com/keegan/product/domain/repository"
)

// 下一步写结构体
type IProductDataService interface {
	AddProduct(*model.Product)(int64,error)
	DeleteProduct(int64)error
	UpdateProduct(*model.Product)error
	FindProductByID(int64)(*model.Product,error)
	FindAllProduct() ([]model.Product,error)
}

// 创建
func NewProductDataService(productRepository repository.IProductRepository) IProductDataService{
	return &ProductDataService{ productRepository }
}

type ProductDataService struct {
	ProductRepository repository.IProductRepository
}

// 插入
func (u *ProductDataService)AddProduct(product *model.Product)(int64,error){
	return u.ProductRepository.CreateProduct(product)
}

// 删除
func (u *ProductDataService)DeleteProduct(productID int64)error{
	return u.ProductRepository.DeleteProductByID(productID)
}

// 更新
func (u *ProductDataService)UpdateProduct(product *model.Product)error{
	return u.ProductRepository.UpdateProduct(product)
}

// 查找
func (u *ProductDataService)FindProductByID(productID int64)(*model.Product,error){
	return u.ProductRepository.FindProductByID(productID)
}

// 查找
func (u *ProductDataService)FindAllProduct()([]model.Product,error) {
	return u.ProductRepository.FindAll()
}
/*一切都是为了解耦*/
