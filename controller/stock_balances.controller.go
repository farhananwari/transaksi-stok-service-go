package controller

import (
	"net/http"
	"transaksi-stok-service-go/dto"
	"transaksi-stok-service-go/services"

	"github.com/gin-gonic/gin"
)

type ImplStokBalancesController struct {
	service services.StokBalancesService
}

func NewStokBalancesController(service services.StokBalancesService) *ImplStokBalancesController {
	return &ImplStokBalancesController{service: service}
}

func (c *ImplStokBalancesController) GetBarangWithStok(ctx *gin.Context) {
	locationId := ctx.Param("locationID")

	if locationId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "locationID is required",
		})
		return
	}

	result, err := c.service.GetBarangWithStok(locationId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (c *ImplStokBalancesController) CreateStokBalance(ctx *gin.Context) {
	var input dto.CreateBarangStokDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.service.CreateStokBalance(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
