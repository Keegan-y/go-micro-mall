package handler

import (
	"context"
	"git.imooc.com/keegan/common"
	"git.imooc.com/keegan/order/domain/model"
	"git.imooc.com/keegan/order/domain/service"
	. "git.imooc.com/keegan/order/proto/order"
)
type Order struct{
	OrderDataService service.IOrderDataService
}
//根据订单ID查询订单
func (o *Order) GetOrderByID(ctx context.Context,request *OrderID,response *OrderInfo) error {
	order,err:=o.OrderDataService.FindOrderByID(request.OrderId)
	if err !=nil {
		return err
	}
	if err:= common.SwapTo(order,response);err !=nil {
		return err
	}
	return nil
}
//查找所有订单
func (o *Order) GetAllOrder(ctx context.Context, request *AllOrderRequest, response *AllOrder) error {
	orderAll,err := o.OrderDataService.FindAllOrder()
	if err !=nil {
		return err
	}

	for _,v:=range orderAll{
		order := &OrderInfo{}
		if err:=common.SwapTo(v,order);err!=nil {
			return err
		}
		response.OrderInfo = append(response.OrderInfo, order)
	}
	return nil
}
//创建订单
func (o *Order) CreateOrder(ctx context.Context,request *OrderInfo,response *OrderID) error {
	orderAdd := &model.Order{}
	if err:=common.SwapTo(request,orderAdd);err!=nil{
		return err
	}
	orderID,err:=o.OrderDataService.AddOrder(orderAdd)
	if err !=nil {
		return err
	}
	response.OrderId = orderID
	return nil
}

//删除订单
func (o *Order) DeleteOrderByID(ctx context.Context, request *OrderID, response *Response) error {
	if err:=o.OrderDataService.DeleteOrder(request.OrderId);err!=nil {
		return err
	}
	response.Msg = "删除成功"
	return  nil
}

//更新订单支付状态
func (o *Order) UpdateOrderPayStatus(ctx context.Context, request *PayStatus,response *Response) error {
	if err:=o.OrderDataService.UpdatePayStatus(request.OrderId, request.PayStatus);err!=nil {
		return err
	}
	response.Msg = "支付状态更新成功"
	return nil
}

//更新发货状态
func (o *Order) UpdateOrderShipStatus(ctx context.Context, request *ShipStatus, response *Response) error {
	if err:=o.OrderDataService.UpdateShipStatus(request.OrderId,request.ShipStatus);err!=nil{
		return err
	}
	response.Msg = "发货状态更新成功"
	return nil
}
//更新订单状态
func (o *Order) UpdateOrder(ctx context.Context, request *OrderInfo, response *Response) error {
	order := &model.Order{}
	if err:=common.SwapTo(request,order);err!=nil{
		return err
	}
	if err:=o.OrderDataService.UpdateOrder(order);err!=nil {
		return err
	}
	response.Msg = "订单更新成功"
	return nil
}