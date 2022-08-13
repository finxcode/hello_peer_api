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
// @Success 0 {object} models.SquareSettings
// @Router /settings/getSquareSetting [get]
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
// @Success 0
// @Router /settings/setSquareSetting [post]
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
// @Success 0 {object} models.RecommendSetting
// @Router /settings/getRecommendSetting [get]
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
// @Success 0
// @Router /settings/setRecommendSetting [post]
func SetRecommendSettings(c *gin.Context) {
	//设置用户推荐设置
}
