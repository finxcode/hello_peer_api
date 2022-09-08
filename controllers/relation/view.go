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
