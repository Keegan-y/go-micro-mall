package model


type Product struct{
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	ProductName string `json:"product_name"`
	ProductSku string `gorm:"unique_index:not_null" json:"product_sku"`
	ProductPrice float64 `json:"product_price"`
	ProductDescription string `json:"product_description"`
	ProductImage []ProductImage `gorm:"ForeignKey:ImageProductID" json:"product_image"`
	ProductSize []ProductSize `gorm:"ForeignKey:SizeProductID" json:"product_size"`
	ProductSeo ProductSeo `gorm:"ForeignKey:SeoProductID" json:"product_seo"`
}