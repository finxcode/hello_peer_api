package settings

// GetUserSettings
// @Summary 用户可个人设置
// @Description 可获取用户个人设置 Verified 用户实名 0-未认证 1-已认证 PetVerified 宠物实名 0-未认证 1-已认证 phone="" 未绑定
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

// GetUserPhoneNumber
// @Summary 可通过微信接口获取授权用户手机号
// @Description 可通过微信接口获取授权用户手机号
// @ID get_user_phone_number
// @Tags Settings
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @param code query string true "code"
// @Success 0 {object} models.UserPhoneNumber
// @Router /user/setting/getPhoneNumber [get]
func GetUserPhoneNumber() {
	//通过微信接口获取用户手机号
}
