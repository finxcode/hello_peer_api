package models

import "github.com/gin-gonic/gin"

type PetResponse struct {
	Pet_Name    string   `json:"pet_Name" example:"狗狗2022"`
	Sex         string   `json:"sex" example:"MM"`
	Birthday    string   `json:"birthday" example:"2021-10-12"`
	Weight      float32  `json:"weight" example:"3.2"`
	Description string   `json:"description" example:"这是一个好宠物"`
	Images      []string `json:"images" example:"img1.jpg,img2.jpg"`
}

type PetRequest struct {
	Pet_Name    string  `json:"pet_Name" example:"狗狗2022"`
	Sex         string  `json:"sex" example:"MM"`
	Birthday    string  `json:"birthday" example:"2021-10-12"`
	Weight      float32 `json:"weight" example:"3.2"`
	Description string  `json:"description" example:"这是一个好宠物"`
}

// GetPetDetails
// @Summary 获取宠物详情
// @Description 可获取宠物详情
// @ID get_pet_details
// @Tags Pet
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 200 {object} PetResponse
// @Router /user/pet/getPetDetails [get]
func GetPetDetails(c *gin.Context) {
	//获取宠物详情
}

// SetPetDetails
// @Summary 设置宠物详情
// @Description 用户可设置宠物详情
// @ID set_pet_details
// @Tags Pet
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Param user body PetRequest true "宠物详情"
// @Success 200
// @Router /user/pet/setPetDetails [post]
func SetPetDetails(c *gin.Context) {
	// 更新宠物详情
}
