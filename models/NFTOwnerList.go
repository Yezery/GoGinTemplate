package models

import "github.com/google/uuid"

type NFTOwnerList struct {
	Id           uuid.UUID `gorm:"column:Id"`
	NFTName      string    `gorm:"column:nftName" json:"nftName"`
	IsActive     int       `gorm:"column:isActive" json:"isActive"`
	Price        int64     `gorm:"column:price" json:"price"`
	Type         string    `gorm:"column:type" json:"type"`
	TokenId      uint64    `gorm:"column:tokenId" json:"tokenId"`
	OwnerAddress string    `gorm:"column:ownerAddress" json:"ownerAddress"`
	NFTAddress   string    `gorm:"column:nftAddress" json:"nftAddress"`
	IpfsPath     string    `gorm:"column:ipfsPath" json:"ipfsPath"`
}
