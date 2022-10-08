package settings

// GetUserSettings
// @Summary 用户可个人设置
// @Description 可获取用户个人设置 UserVerified 用户实名 0-未认证 1-已认证 PetVerified 宠物实名 0-未认证 1-已认证 phone="" 未绑定
// @ID get_user_settings
// @Tags Settings
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {object} models.UserSetting
// @Router /user/setting/getUserSettings [get]
func GetUserSettings() {
	//用户可获取个人设置
}
