package controllers

import "github.com/gin-gonic/gin"

// AutoLogin
// @Summary 未授权用户登录
// @Description 通过传递微信登录code换取用户token
// @ID autologin
// @Tags Auth
// @Accept application/json
// @Produce application/json
// @Param code body models.AutoLogin true "微信登录code"
// @Success 0 {object} models.Token
// @Router /auth/autologin [post]
func AutoLogin(c *gin.Context) {
	//未授权用户登录
}

// AuthLogin
// @Summary 授权用户登录
// @Description 通过传递微信登录code以及getUserProfile返回字段换取用户token
// @ID authlogin
// @Tags Auth
// @Accept application/json
// @Produce application/json
// @Param code body models.UserProfileForm true "微信登录code以及getUserProfile返回字段"
// @Success 0 {object} models.Token
// @Router /auth/authlogin [post]
func AuthLogin(c *gin.Context) {
	//授权用户登录
}
