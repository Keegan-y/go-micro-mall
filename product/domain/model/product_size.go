package model

type ProductSize struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	SizeName string `json:"size_name"`
	SizeCode string `gorm:"unique_index;not_null" json:"size_code"`
	SizeProductID int64 `json:"size_product_id"`
}
