package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Driver string
	DSN    string
}

type SaleList struct {
	SaleId      int64  `gorm:"primaryKey;column:saleId"`
	Seller      string `gorm:"column:seller"`
	NftContract string `gorm:"column:nftContract"`
	Tokenid     int64  `gorm:"column:tokenid"`
	Price       int64  `gorm:"column:price"`
	IsActive    bool   `gorm:"column:isActive"`
}

func main() {
	router := gin.Default()
	dataBytes, err := os.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("读取配置文件失败：", err)
		return
	}
	// fmt.Println("yaml 文件的内容: \n", string(dataBytes))
	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	// fmt.Printf("config → %+v\n", config)

	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(config.Database.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 包含数据模型的切片作为参数，并自动检查、创建和修改数据库表结构以匹配模型定义。

	_ = db.AutoMigrate(&SaleList{})

	// 查询用户表中的全部记录
	var SaleLists []SaleList
	result := db.Find(&SaleLists)
	if result.Error != nil {
		// 处理错误逻辑
	}
	for _, sale := range SaleLists {
		fmt.Printf("saleId: %d, seller: %s, nftContract: %s\n", sale.SaleId, sale.Seller, sale.NftContract)
	}
	
	router.Run(":" + config.Server.Port)
}
