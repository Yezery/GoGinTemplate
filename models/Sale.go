package models

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	SaleId      int64  `gorm:"primaryKey;column:sale_id" json:"saleId"`
	Type        string `gorm:"column:type" json:"type"`
	Seller      string `gorm:"column:seller" json:"seller"`
	NftContract string `gorm:"column:nft_contract" json:"nftContract"`
	Tokenid     int64  `gorm:"column:token_id" json:"tokenid"`
	Price       int64  `gorm:"column:price" json:"price"`
	IsActive    bool   `gorm:"column:is_active" json:"isActive"`
}
