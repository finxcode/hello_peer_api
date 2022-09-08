package relation

import "github.com/gin-gonic/gin"

// UpdateAllNewFocusStatus
// @Summary 设置用户新增关注已阅
// @Description 可设置用户新增关注已阅
// @ID update_all_new_focus_status
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0
// @Router /user/relation/updateAllNewFocusStatus [post]
func UpdateAllNewFocusStatus(c *gin.Context) {
	//设置用户新增看过已阅
}
