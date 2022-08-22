package handler

import (
	"context"
	"git.imooc.com/keegan/category/common"
	"git.imooc.com/keegan/category/domain/model"
	"git.imooc.com/keegan/category/domain/service"
	category "git.imooc.com/keegan/category/proto/category"
	"github.com/prometheus/common/log"
)

type Category struct{
	CategoryDataService service.ICategoryDataService
}

// 提供创建分类的服务
func (c *Category)CreateCategory(ctx context.Context,request *category.CategoryRequest,response *category.CreateCategoryResponse) error {
	category := &model.Category{}
	// 赋值 common/swap.go
	err := common.SwapTo(request,category)
	if err != nil {
		return err
	}
	categoryId,err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}
	response.Message = "分类添加成功"
	response.CategoryId = categoryId

	return nil

}

// 提供分类更新服务
func (c *Category)UpdateCategory(ctx context.Context,request *category.CategoryRequest,response *category.UpdateCategoryResponse) error{
	category := &model.Category{}
	err := common.SwapTo(response,category)
	if err != nil{
		return err
	}
	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil{
		return err
	}
	response.Message = "分类更新成功"

	return nil
}

// 提供分类删除服务
func (c *Category) DeleteCategory(ctx context.Context,request *category.DeleteCategoryRequest,response *category.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(request.CategoryId)
	if err != nil {
		return nil
	}
	response.Message = "删除成功"

	return nil
}

// 根据分类名称来查找分类
func (c *Category) FindCategoryByName(ctx context.Context,request *category.FindByNameRequest,response *category.CategoryResponse) error {
	category,err := c.CategoryDataService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return err
	}

	return common.SwapTo(category,response)
}

// 根据分类ID来查找分类
func (c *Category) FindCategoryByID(ctx context.Context,request *category.FindByIdRequest,response *category.CategoryResponse) error {
	category,err := c.CategoryDataService.FindCategoryByID(request.CategoryId)

	if err != nil {
		return err
	}

	return common.SwapTo(category,response)
}

func (c *Category) FindCategoryByLevel(ctx context.Context,request *category.FindByLevelRequest,response *category.FindAllResponse) error {
	categorySlice,err:=c.CategoryDataService.FindCategoryByLevel(request.Level)
	if err !=nil {
		return err
	}
	categoryToResponse(categorySlice,response)
	return nil
}

func (c *Category) FindCategoryByParent(ctx context.Context,request *category.FindByParentRequest,response *category.FindAllResponse) error {
	categorySlice,err := c.CategoryDataService.FindCategoryByParent(request.ParentId)
	if err != nil {
		return err
	}

	categoryToResponse(categorySlice,response)
	return nil
}

func (c *Category) FindAllCategory(ctx context.Context,request *category.FindAllRequest,response *category.FindAllResponse) error {
	categorySlice,err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return nil
	}
	categoryToResponse(categorySlice,response)

	return nil
}

func categoryToResponse(categorySlice []model.Category,response *category.FindAllResponse){
	for _,cg := range categorySlice {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(cg,cr)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category,cr)
	}
}
