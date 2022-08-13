package controllers

import "github.com/gin-gonic/gin"

// GetUserIMSig
// @Summary 获取用户腾讯IM签名
// @Description 可获取用户腾讯IM签名
// @ID get_user_im_sig
// @Tags Tencent
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 200 {object} models.UserIMSig
// @Router /user/tencent/getUserIMSig [get]
func GetUserIMSig(c *gin.Context) {
	// 获取用户腾讯IM签名
}
