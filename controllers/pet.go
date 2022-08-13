package controllers

import "github.com/gin-gonic/gin"

// DeletePetImage
// @Summary 删除宠物图片
// @Description 删除宠物图片
// @ID delete_pet_image
// @Tags Pet
// @Security x-token
// @param x-token header string true "Authorization"
// @param filename query string true "文件名"
// @Success 0
// @Router /user/pet/deletePetImage [post]
func DeletePetImage(c *gin.Context) {
	// 删除宠物图片
}

// SetPetImage
// @Summary 设置宠物图片
// @Description 用户可上传宠物图片
// @ID set_pet_image
// @Tags Pet
// @Accept mpfd
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @param content formData file true "宠物图片"
// @Success 0
// @Router /user/pet/upload/setPetImage [post]
func SetPetImage(c *gin.Context) {
	//设置宠物图片
}

// InitPet
// @Summary 初始化宠物
// @Description 用户可初始化宠物
// @ID init_pet
// @Tags Pet
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0
// @Router /user/pet/initPet [post]
func InitPet(c *gin.Context) {
	// 初始化宠物
}
