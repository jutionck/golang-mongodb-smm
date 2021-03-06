package controller

import (
	"enigmacamp.com/golang-with-mongodb/model"
	"enigmacamp.com/golang-with-mongodb/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ProductController struct {
	router         *gin.Engine
	productUseCase usecase.ProductRegistrationUseCase
}

func (pc *ProductController) registerNewProduct(ctx *gin.Context) {
	var newProduct model.Product
	err := ctx.ShouldBindJSON(&newProduct)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = pc.productUseCase.Register(&newProduct)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
		"data":    newProduct,
	})
}

func NewProductController(router *gin.Engine, productUseCase usecase.ProductRegistrationUseCase) *ProductController {
	controller := ProductController{
		router:         router,
		productUseCase: productUseCase,
	}
	router.POST("/product", controller.registerNewProduct)
	return &controller
}
