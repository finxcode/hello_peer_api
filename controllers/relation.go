package controllers

import "github.com/gin-gonic/gin"

// SetFocusOn
// @Summary 设置两个用户之间的关注
// @Description 可根据用户ID，目标用户ID，关注状态设置用户之间的关注。 0 - 取消关注， 1 - 新关注， 2 - 已关注（用户已阅）
// @ID set_focus_on
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {object} models.FocusRequest
// @Router /user/relation/setFocusOn [post]
func SetFocusOn(c *gin.Context) {
	//设置关注状态
}
