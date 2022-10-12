package sys

// GetUserTerms
// @Summary 可获取用户服务协议
// @Description 可获取用户协议
// @ID get_user_terms
// @Tags Settings
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {object} models.Agreement
// @Router /sys/getUserTerms [get]
func GetUserTerms() {
	//可获取用户服务协议
}

// GetPrivacyPolicy
// @Summary 可获取隐私政策
// @Description 可获取隐私政策
// @ID get_privacy_policy
// @Tags Settings
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {object} models.Agreement
// @Router /sys/getPrivacyPolicy [get]
func GetPrivacyPolicy() {
	//可获取隐私政策
}
