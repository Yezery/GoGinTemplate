package controllers

import (
	"net/http"

	"example.com/m/models"
	"example.com/m/repositories"
	"example.com/m/utils"
	"github.com/gin-gonic/gin"
)

type SaleController struct{}

var Sale models.Sale

// 查询上架表中的全部记录
func (SC *SaleController) GetSaleList(c *gin.Context) {
	var Sales []models.Sale
	db := repositories.GetDb(c)
	result := db.Find(&Sales)
	if result.Error != nil {
		utils.SendResponse(c.Writer, http.StatusInternalServerError, result.Error)
		panic(result.Error)
	}
	utils.SendResponse(c.Writer, http.StatusOK, Sales)
}

// 创建一条新记录
func (SC *SaleController) CreateSale(c *gin.Context) {
	repositories.MakeDbModel[models.Sale](&Sale, c)
	if err := c.BindJSON(&Sale); err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
		return
	}
	db := repositories.GetDb(c)
	result := db.Create(&Sale)
	if result.Error != nil {
		utils.SendResponse(c.Writer, http.StatusInternalServerError, result.Error)
		panic(result.Error)
	}
	utils.SendResponse(c.Writer, http.StatusOK, Sale)
}
