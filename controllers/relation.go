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

// getFans
// @Summary 获取用户我的粉丝列表
// @Description 可获取用户的粉丝列表。 status说明， 0 - 关注ta， 1 - 已互关
// @ID get_fans
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {array} models.Fan
// @Router /user/relation/getFans [get]
func getFans(c *gin.Context) {
	//获取用户我的粉丝
}

// getFansToOthers
// @Summary 获取用户关注的人的列表
// @Description 可获取用户关注的人的列表。 status说明， 0 - 已关注， 1 - 已互关
// @ID get_fans_to_others
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {array} models.Fan
// @Router /user/relation/getFansToOthers [get]
func getFansToOthers(c *gin.Context) {
	//获取用户关注的人的列表
}
