package model

type OrderDetail struct {
	ID int64 `grom:"primary_key;not_null;auto_increment" json:"id"`
	ProductID int64 `json:"product_id"`
	ProductNum int64 `json:"product_num"`
	ProductSizeID int64 `json:"product_size_id"`
	ProductPrice float64 `json:"product_price"`
	OrderID int64 `json:"order_id"`
}