package repository

import (
	"git.imooc.com/keegan/category/domain/model"
	"github.com/jinzhu/gorm"
)
type ICategoryRepository interface {
	InitTable() error
	FindCategoryByID(int64) (*model.Category,error)
	CreateCategory(*model.Category) (int64,error)
	DeleteCategoryByID(int64) error
	UpdateCategory(*model.Category) error
	FindAll() ([]model.Category,error)
	FindCategoryByName(string) (*model.Category,error)
	FindCategoryByLevel(uint32) ([]model.Category,error)
	FindCategoryByParent(int64) ([]model.Category,error)
}

// 创建 categoryRepository
type CategoryRepository struct {
	mysqlDb *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb:db}
}

// 初始化表
func (u *CategoryRepository)InitTable() error {
	return u.mysqlDb.CreateTable(&model.Category{}).Error
}


// 创建Category信息
func (u *CategoryRepository) CreateCategory(category *model.Category) (int64,error){
	return category.ID,u.mysqlDb.Create(category).Error
}

// 根据ID删除Category信息
func (u *CategoryRepository) DeleteCategoryByID(categoryID int64) error{
	return u.mysqlDb.Where("id=?",categoryID).Delete(&model.Category{}).Error
}

// 更新Category信息
func (u *CategoryRepository) UpdateCategory(category *model.Category) error {
	return u.mysqlDb.Model(category).Update(category).Error
}

// 根据ID查找Category信息
func (u *CategoryRepository)FindCategoryByID(categoryID int64) (category *model.Category,err error){
	category = &model.Category{}
	return category,u.mysqlDb.First(category,categoryID).Error
}

// 根据分类的名称来查找
func (u *CategoryRepository) FindCategoryByName(categoryName string) (category *model.Category,err error) {
	category = &model.Category{}
	return category,u.mysqlDb.Where("category_name=?",categoryName).Find(category).Error
}

// 根据分类等级来查找
func (u *CategoryRepository) FindCategoryByLevel (level uint32) (categorySlice []model.Category,err error) {
	return categorySlice,u.mysqlDb.Where("category_level=?",level).Find(categorySlice).Error
}

// 查找父级分类
func (u *CategoryRepository) FindCategoryByParent(parent int64)(categorySlice []model.Category,err error) {
	return categorySlice,u.mysqlDb.Where("category_parent=?",parent).Find(categorySlice).Error
}

// 获取结果集
func (u *CategoryRepository) FindAll()(categoryAll []model.Category,err error){
	return categoryAll,u.mysqlDb.Find(&categoryAll).Error
}