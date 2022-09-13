package relation

import "github.com/gin-gonic/gin"

// UpdateAllNewViewStatus
// @Summary 设置用户新增看过已阅
// @Description 可设置用户新增看过已阅
// @ID update_all_new_view_status
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0
// @Router /user/relation/updateAllNewViewStatus [post]
func UpdateAllNewViewStatus(c *gin.Context) {
	//设置用户新增看过已阅
}

// GetViewList
// @Summary 获取用户关注列表
// @Description 可获取用户关注列表 status说明： 0 - 新关注（未揭秘） 1 - 已阅关注（未揭秘） 2 - 已揭秘
// @ID get_view_list
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {array} models.View
// @Router /user/relation/getViewList [get]
func GetViewList(c *gin.Context) {
	//获取用户关注列表
}
