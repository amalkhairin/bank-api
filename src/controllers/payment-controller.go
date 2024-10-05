package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"bank-api/src/services"
)

type PaymentController struct {
	paymentService *services.PaymentService
}

func NewPaymentController(paymentService *services.PaymentService) *PaymentController {
	return &PaymentController{paymentService: paymentService}
}

func (c *PaymentController) CreatePayment(ctx *gin.Context) {
	userID := ctx.GetString("userID")

	var paymentRequest struct {
		Amount      float64 `json:"amount" binding:"required"`
		Description string  `json:"description" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&paymentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.paymentService.CreatePayment(userID, paymentRequest.Amount, paymentRequest.Description)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Payment successful"})
}

func (c *PaymentController) GetUserPayments(ctx *gin.Context) {
	userID := ctx.GetString("userID")

	payments, err := c.paymentService.GetUserPayments(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"payments": payments})
}
