package controllers

import (
	"BCScanner/config"
	"BCScanner/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type QRGenerateController struct {
	QRGenerateUseCase domain.QRGenerateUseCase
	Env               *config.Env
}

func (ctrl *QRGenerateController) CreateQR(c *gin.Context) {
	var QrGenerateRequest domain.QRGenerateRequest
	//var QrGenerateResponse domain.QRGenerateResponse
	if err := c.ShouldBind(&QrGenerateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	QrGenerateResponse, err := ctrl.QRGenerateUseCase.CreateQR(c, &QrGenerateRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, QrGenerateResponse)
}

func (ctrl *QRGenerateController) FetchQR(c *gin.Context) {
	// Call the use case to fetch data
	//results, err := ctrl.QRGenerateUseCase.Fetch(c.Request.Context())
	results, err := ctrl.QRGenerateUseCase.Fetch(c)
	if err != nil {
		// Return error response
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return successful response with results
	c.JSON(http.StatusOK, results)
}

func (ctrl *QRGenerateController) GetByID(c *gin.Context) {
	id := c.Query("id")
	QRGenerate, err := ctrl.QRGenerateUseCase.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, QRGenerate)
}
