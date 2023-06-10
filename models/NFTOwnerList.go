package models

import "gorm.io/gorm"

type NFTOwnerList struct {
	gorm.Model
	OwnerAddress string `gorm:"column:ownerAddress" json:"ownerAddress"`
	NFTAddress   string `gorm:"column:nftAddress" json:"nftAddress"`
	IpfsPath     string `gorm:"column:ipfsPath" json:"ipfsPath"`
}
