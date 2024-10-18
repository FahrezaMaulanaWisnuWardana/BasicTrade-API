package controllers

import (
	"BasicTrade-API/database"
	"BasicTrade-API/helpers"
	"BasicTrade-API/models/entity"
	"BasicTrade-API/models/request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetVariants(ctx *gin.Context) {
	db := database.GetDB()
	var Variant []entity.Variant
	search := ctx.Query("search")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	result := db.Debug().Model(&entity.Variant{}).Where("variant_name LIKE ?", "%"+search+"%").Offset((page - 1) * limit).Limit(limit).Find(&Variant).Error
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
		Data:       Variant,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": pagination,
	})
}

func AddVariants(ctx *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(ctx)
	var VariantRequest request.VariantRequest
	if contentType == appJson {
		ctx.ShouldBindJSON(&VariantRequest)
	} else {
		ctx.ShouldBind(&VariantRequest)
	}
	var product entity.Product
	result := db.Where("uuid = ?", VariantRequest.ProductUUID).First(&product)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": result.Error,
		})
		return
	}
	newUUID := uuid.New()
	Variant := entity.Variant{
		UUID:        newUUID.String(),
		Quantity:    VariantRequest.Quantity,
		VariantName: VariantRequest.VariantName,
		ProductID:   product.ID,
	}
	err := db.Debug().Create(&Variant).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Variant,
	})
}

func UpdateVariants(ctx *gin.Context) {
	db := database.GetDB()
	VariantUUID := ctx.Param("uuid")

	var VariantReq request.VariantRequest

	contentType := helpers.GetContentType(ctx)

	if contentType == appJson {
		ctx.ShouldBindJSON(&VariantReq)
	} else {
		ctx.ShouldBind(&VariantReq)
	}

	// Retrieve existing book from the database
	var GetVariant entity.Variant
	if err := db.Where("uuid = ?", VariantUUID).First(&GetVariant).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	GetVariant.VariantName = VariantReq.VariantName
	GetVariant.Quantity = VariantReq.Quantity
	db.Save(&GetVariant)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data Updated Successfully",
	})
}

func DeleteVariants(ctx *gin.Context) {
	db := database.GetDB()
	VariantUUID := ctx.Param("uuid")
	Variant := entity.Variant{}
	err := db.Debug().Where("uuid = ?", VariantUUID).Delete(&Variant).Error
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

func VariantsByUUID(ctx *gin.Context) {
	db := database.GetDB()
	VariantUUID := ctx.Param("uuid")
	var Variant entity.Variant
	err := db.Debug().Model(&entity.Variant{}).Where("uuid = ?", VariantUUID).Find(&Variant).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": Variant,
	})
}
