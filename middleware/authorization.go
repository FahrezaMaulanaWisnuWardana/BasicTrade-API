package middleware

import (
	"BasicTrade-API/database"
	"BasicTrade-API/models/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		ProductUUID := ctx.Param("uuid")

		adminData := ctx.MustGet("adminData").(jwt.MapClaims)
		adminID := uint(adminData["id"].(float64))

		var getProduct entity.Product
		err := db.Debug().Select("admin_id").Where("uuid = ?", ProductUUID).First(&getProduct).Error
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Product Not Found",
			})
			return
		}

		if getProduct.AdminID != adminID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
	}
}
