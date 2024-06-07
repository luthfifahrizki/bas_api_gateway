package handler

import (
	"api_gateway/model"
	"api_gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionInterface interface {
	TransferBank(*gin.Context)
	CreateDataTransaksi(*gin.Context)
}

type transactionImplement struct{}

func NewTransaction() TransactionInterface {
	return &transactionImplement{}
}

type BodyPayloadTransaction struct{}

func (b *transactionImplement) TransferBank(g *gin.Context) {

	bodyPayloadTxn := BodyPayloadTransaction{}
	err := g.BindJSON(&bodyPayloadTxn)

	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this API rest for later",
	})
}

func (a *transactionImplement) CreateDataTransaksi(g *gin.Context) {
	bodyPayload := model.Transaction{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Create(&bodyPayload)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Create Data Transaksi successfully",
		"data":    bodyPayload,
	})
}
