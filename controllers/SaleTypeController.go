package controllers

import (
	"net/http"

	"example.com/m/models"
	"example.com/m/repositories"
	"github.com/gin-gonic/gin"
)

type SaleTypeController struct{}

var TypeModel models.SaleType

// 查询上架表中的全部记录
func (STC *SaleTypeController) GetTypeList(c *gin.Context) {
	var Type []models.SaleType
	repositories.MakeDbModel[models.SaleType](&TypeModel, c)
	db := repositories.GetDb(c)
	result := db.Find(&Type)
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, result.Error)
		panic(result.Error)
	}
	c.JSON(http.StatusOK, Type)
}
