package controllers

import (
	"net/http"

	"example.com/m/models"
	"example.com/m/repositories"
	"example.com/m/utils"
	"github.com/gin-gonic/gin"
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
	result := db.Create(&NFTOwnerList)
	if result.Error != nil {
		utils.SendResponse(c.Writer, http.StatusInternalServerError, result.Error)
		panic(result.Error)
	}
	utils.SendResponse(c.Writer, http.StatusOK, "success")
}
