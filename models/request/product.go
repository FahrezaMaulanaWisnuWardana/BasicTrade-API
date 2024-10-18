package request

import "mime/multipart"

type ProductRequest struct {
	Name     string                `form:"name" binding:"required"`
	ImageUrl *multipart.FileHeader `form:"file" binding:"required"`
}
