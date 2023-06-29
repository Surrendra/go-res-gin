package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-res-gin/helpers"
	"github.com/go-res-gin/middlewares"
	"github.com/go-res-gin/models"
	"github.com/google/uuid"
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

	transactionModel := models.Transaction{
		Code:          uuid.New().String(),
		CustomerName:  transactionRequest.CustomerName,
		CreatedUserID: authUser.Id,
	}

	models.DB.Create(&transactionModel)
	var grandTotal int64 = 0
	for _, item := range transactionRequest.Items {
		subTotal := item.Qty * item.ProductPrice
		grandTotal += subTotal

		transactionProduct := models.TransactionProduct{
			TransactionId: transactionModel.Id,
			ProductId:     item.ProductID,
			ProductName:   item.ProductName,
			Quantity:      item.Qty,
			ProductPrice:  item.ProductPrice,
			SubTotal:      subTotal,
		}
		models.DB.Create(&transactionProduct)
		// index is the index where we are
		// element is the element from someSlice for where we are
	}

	// update total price from grand total to transactionModel
	transactionModel.TotalPrice = grandTotal
	models.DB.Save(&transactionModel)
	resHelper.ResponseSuccess(c, transactionModel, "Success create transaction")
}
