package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-res-gin/helpers"
	"github.com/go-res-gin/middlewares"
	"github.com/go-res-gin/models"
	"github.com/google/uuid"
	"time"
)

var ResHelper = helpers.NewResHelper()

type TransactionHandler struct {
	JwtMiddleware middlewares.JwtMiddleware
}

func NewTransactionHandler(JwtMiddleware middlewares.JwtMiddleware) *TransactionHandler {
	return &TransactionHandler{
		JwtMiddleware: JwtMiddleware,
	}
}

func (h TransactionHandler) GetData(c *gin.Context) {
	var transaction models.Transaction
	models.DB.Find(&transaction)
	resHelper.ResponseSuccess(c, transaction, "")
}

func (h TransactionHandler) Create(c *gin.Context) {
	type TransactionRequest struct {
		CustomerName string `json:"customer_name"`
		Items        []struct {
			ProductID    int64  `json:"product_id"`
			ProductName  string `json:"product_name"`
			ProductPrice int64  `json:"productPrice"`
			Qty          int64  `json:"Qty"`
		} `json:"items"`
	}
	var transactionRequest TransactionRequest
	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		fmt.Println("Something wrong when binding the request")
		resHelper.ResponseValidationError(c, err, "Something wrong with the request")
		return
	}

	authUser, _ := h.JwtMiddleware.GetAuthUser(c)

	Transaction := models.Transaction{
		Code:            uuid.New().String(),
		CustomerName:    transactionRequest.CustomerName,
		TransactionDate: time.Now().Format("2006-01-02"),
		CreatedUserID:   authUser.Id,
	}

	models.DB.Create(&Transaction)
	var grandTotal int64 = 0
	var product models.Product
	for _, item := range transactionRequest.Items {
		subTotal := item.Qty * item.ProductPrice
		grandTotal += subTotal

		// clean product models to re select
		product = models.Product{}
		models.DB.Where("id = ?", item.ProductID).Find(&product)

		transactionProduct := models.TransactionProduct{
			TransactionId: Transaction.Id,
			ProductId:     item.ProductID,
			ProductName:   product.Name,
			Quantity:      item.Qty,
			ProductPrice:  item.ProductPrice,
			SubTotal:      subTotal,
		}
		models.DB.Create(&transactionProduct)
	}

	// update total price from grand total to transactionModel
	Transaction.TotalPrice = grandTotal
	models.DB.Save(&Transaction)

	// select transaction with CreatedUser and TransactionProduct relations
	models.DB.Preload("CreatedUser").Preload("TransactionProducts").Find(&Transaction, Transaction.Id)

	resHelper.ResponseSuccess(c, Transaction, "Success create transaction")
}
