package dto

import "table_management/entity"

type OrderRequest struct {
	CustomerName string
	TableId      string
	Orders       []OrderDetailRequest
}

type OrderDetailRequest struct {
	MenuId string
	Qty    int
	Price  int
}

func (or *OrderRequest) ToOrderEntity() (customerOrder *entity.CustomerOrder) {
	var orderDetails []entity.CustomerOrderDetail
	for _, od := range or.Orders {
		orderDetail := entity.CustomerOrderDetail{
			MenuID: od.MenuId,
			Qty:    od.Qty,
			Price:  od.Price,
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	customerOrder = new(entity.CustomerOrder)
	customerOrder.CustomerName = or.CustomerName
	customerOrder.OrderDetails = orderDetails
	return
}
