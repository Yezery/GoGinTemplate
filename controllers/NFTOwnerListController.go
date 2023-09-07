package controllers

import (
	"net/http"

	"example.com/m/models"
	"example.com/m/repositories"
	"example.com/m/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NFTOwnerListController struct{}

var NFTOwnerList models.NFTOwnerList

// 创建一条新记录
func (NFTLC *NFTOwnerListController) CreateNFTInf(c *gin.Context) {
	repositories.MakeDbModel[models.NFTOwnerList](&NFTOwnerList, c)
	if err := c.BindJSON(&NFTOwnerList); err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, err)
		panic(err)
	}
	db := repositories.GetDb(c)
	// 生成UUID
	id := uuid.New()
	// 将UUID分配给Id字段
	NFTOwnerList.Id = id
	result := db.Create(&NFTOwnerList)
	if result.Error != nil {
		utils.SendResponse(c.Writer, http.StatusInternalServerError, result.Error)
		panic(result.Error)
	}
	utils.SendResponse(c.Writer, http.StatusOK, "success")
}  

// 查询属于个人的NFT
func (NFTLC *NFTOwnerListController) GetOwnerNFTs(c *gin.Context) {
	var nftOwnerList []models.NFTOwnerList
	c.BindJSON(&NFTOwnerList)
	db := repositories.GetDb(c)
	results := db.Where("ownerAddress", NFTOwnerList.OwnerAddress).Find(&nftOwnerList)
	if results.Error != nil {
		utils.SendResponse(c.Writer, http.StatusInternalServerError, results.Error)
		panic(results.Error)
	}
	utils.SendResponse(c.Writer, http.StatusOK, nftOwnerList)
}
