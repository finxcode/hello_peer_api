package controllers

import "github.com/gin-gonic/gin"

// GetSquareSettings
// @Summary 获取用户广场设置
// @Description 可通过用户ID获取用户广场设置
// @ID get_square_settings
// @Tags Settings
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 200 {object} models.SquareSettings
// @Router /settings/getSquareSetting/{user_id} [get]
func GetSquareSettings(c *gin.Context) {
	//获取用户广场设置
}

// SetSquareSettings
// @Summary 设置用户广场设置
// @Description 可设置用户广场设置
// @ID set_square_settings
// @Tags Settings
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Param settings body models.SquareSettings true "设置"
// @Success 200
// @Router /settings/setSquareSetting/{user_id} [post]
func SetSquareSettings(c *gin.Context) {
	//设置用户广场设置
}

// GetRecommendSettings
// @Summary 获取用户推荐设置
// @Description 可通过用户ID获取用户推荐设置
// @ID get_recommend_settings
// @Tags Settings
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 200 {object} models.RecommendSetting
// @Router /settings/getRecommendSetting/{user_id} [get]
func GetRecommendSettings(c *gin.Context) {
	// 获取用户推荐设置
}

// SetRecommendSettings
// @Summary 设置用户广场设置
// @Description 可设置用户广场设置
// @ID set_recommend_settings
// @Tags Settings
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Param settings body models.RecommendSetting true "设置"
// @Success 200
// @Router /settings/setRecommendSetting/{user_id} [post]
func SetRecommendSettings(c *gin.Context) {
	//设置用户推荐设置
}
