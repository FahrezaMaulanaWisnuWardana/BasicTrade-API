package controllers

import (
	"BasicTrade-API/database"
	"BasicTrade-API/helpers"
	"BasicTrade-API/models/entity"
	"BasicTrade-API/models/request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GetProducts(ctx *gin.Context) {
	db := database.GetDB()
	var Products []entity.Product
	search := ctx.Query("search")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	result := db.Debug().Model(&entity.Product{}).Where("name LIKE ?", "%"+search+"%").Preload("Variants").Offset((page - 1) * limit).Limit(limit).Find(&Products).Error
	if result != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	var total int64
	db.Model(&entity.Product{}).Count(&total)
	totalPages := int(total / int64(limit))
	if total%int64(limit) > 0 {
		totalPages++
	}

	pagination := entity.Pagination{
		Page:       page,
		Limit:      limit,
		Total:      int(total),
		TotalPages: totalPages,
		Data:       Products,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": pagination,
	})
}

func AddProduct(ctx *gin.Context) {
	db := database.GetDB()
	adminData := ctx.MustGet("adminData").(jwt.MapClaims)

	var ProductReq request.ProductRequest
	contentType := helpers.GetContentType(ctx)
	if contentType == appJson {
		ctx.ShouldBindJSON(&ProductReq)
	} else {
		ctx.ShouldBind(&ProductReq)
	}
	// Extract the filename without extension
	fileName := helpers.RemoveExtension(ProductReq.ImageUrl.Filename)

	uploadResult, err := helpers.UploadFile(ProductReq.ImageUrl, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUUID := uuid.New()
	Product := entity.Product{
		UUID:     newUUID.String(),
		Name:     ProductReq.Name,
		ImageUrl: uploadResult,
		AdminID:  uint(adminData["id"].(float64)),
	}

	err = db.Debug().Create(&Product).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	var ProductDisplay entity.Product
	result := db.Where("id = ?", Product.ID).Preload("Admin").First(&ProductDisplay)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": result.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": ProductDisplay,
	})
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	ProductUUID := ctx.Param("uuid")

	var ProductReq request.ProductRequest

	contentType := helpers.GetContentType(ctx)

	if contentType == appJson {
		ctx.ShouldBindJSON(&ProductReq)
	} else {
		ctx.ShouldBind(&ProductReq)
	}

	// Retrieve existing book from the database
	var getProduct entity.Product
	if err := db.Where("uuid = ?", ProductUUID).First(&getProduct).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// Extract the filename without extension
	fileName := helpers.RemoveExtension(ProductReq.ImageUrl.Filename)
	uploadResult, err := helpers.UploadFile(ProductReq.ImageUrl, fileName)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	getProduct.Name = ProductReq.Name
	getProduct.ImageUrl = uploadResult
	db.Save(&getProduct)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data Updated Successfully",
	})
}

func DeleteProduct(ctx *gin.Context) {
	db := database.GetDB()
	ProductUUID := ctx.Param("uuid")
	Product := entity.Product{}
	err := db.Debug().Where("uuid = ?", ProductUUID).Delete(&Product).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data Deleted Successfully",
	})
}

func ProductByUUID(ctx *gin.Context) {
	db := database.GetDB()
	ProductUUID := ctx.Param("uuid")
	var Product struct {
		UUID     string
		Name     string
		ImageURL string
	}
	err := db.Debug().Model(&entity.Product{}).Where("uuid = ?", ProductUUID).Find(&Product).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": Product,
	})
}
