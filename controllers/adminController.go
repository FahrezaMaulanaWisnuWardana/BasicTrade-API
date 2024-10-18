package controllers

import (
	"BasicTrade-API/database"
	"BasicTrade-API/helpers"
	"BasicTrade-API/models/entity"
	"BasicTrade-API/models/request"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	appJson = "application/json"
)

func Register(ctx *gin.Context) {
	db := database.GetDB()
	var RequestAdmin request.AdminRequest
	contentType := helpers.GetContentType(ctx)

	if contentType == appJson {
		ctx.ShouldBindJSON(&RequestAdmin)
	} else {
		ctx.ShouldBind(&RequestAdmin)
	}
	// Generate UUID
	newUUID := uuid.New()
	Admin := entity.Admin{
		UUID:     newUUID.String(),
		Email:    RequestAdmin.Email,
		Name:     RequestAdmin.Name,
		Password: RequestAdmin.Password,
	}

	err := db.Debug().Create(&Admin).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Admin,
	})
}
func Login(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	Admin := entity.Admin{}
	var password string

	if contentType == appJson {
		ctx.ShouldBindJSON(&Admin)
	} else {
		ctx.ShouldBind(&Admin)
	}

	password = Admin.Password

	// Get Data Admin
	err := db.Debug().Where("email = ?", Admin.Email).Take(&Admin).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email",
		})
		return
	}
	// Compare Password
	comparePass := helpers.ComparePass([]byte(Admin.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Password",
		})
		return
	}
	token := helpers.GenerateToken(Admin.ID, Admin.Email)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
