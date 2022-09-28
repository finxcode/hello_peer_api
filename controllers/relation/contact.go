package relation

// SendFriendRequest
// @Summary 用户发送认识请求
// @Description 可根据用户ID，给目标用户ID发送认识请求
// @ID send_friend_request
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Param focus body models.ContactRequest true "发送认识请求"
// @Success 0
// @Router /user/relation/sendFriendRequest [post]
func SendFriendRequest() {
	//用户可发送好友申请
}

// ApproveFriendRequest
// @Summary 用户可同意其他用户的认识请求
// @Description 可根据用户ID，同意目标用户ID的认识请求 error_code -1: 申请状态不允许重复发送请求
// @ID approve_friend_request
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Param focus body models.ContactApproveRequest true "同意认识请求"
// @Success 0
// @Router /user/relation/approveFriendRequest [post]
func ApproveFriendRequest() {
	//用户可同意其他用户的认识请求
}

// ReleaseFriendRelation
// @Summary 用户可解除与其他用户的好友关系
// @Description 可根据用户ID，解除目标用户ID的认识关系 error_code -1: 表示用户没有好友关系或者状态不支持解除
// @ID release_friend_relation
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Param focus body models.ContactApproveRequest true "解除认识关系"
// @Success 0
// @Router /user/relation/releaseFriendRelation [post]
func ReleaseFriendRelation() {
	//用户可解除与其他用户的好友关系
}

// GetFriendList
// @Summary 用户可获取好友列表
// @Description 可根据用户ID，获取用户好友列表 宠物类型说明 0-猫 1-狗 用户性别说明 1-男 2-女
// @ID get_friend_list
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {array} models.Friend
// @Router /user/relation/getFriendList [get]
func GetFriendList() {
	//用户可解除与其他用户的好友关系
}

// GetFriendsInSevenDays
// @Summary 用户可获取7日内认识请求
// @Description 可根据用户ID，获取7日内认识请求 state说明： 0-待处理 1-已婉拒 2-过期自动拒绝 3-已同意
// @ID get_friends_in_seven_days
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {array} models.FriendToMeResponse
// @Router /user/relation/getFriendsInSevenDays [get]
func GetFriendsInSevenDays() {
	//用户可获取7日内认识请求
}

// GetFriendsOutOfSevenDays
// @Summary 用户可获取7日前认识请求
// @Description 可根据用户ID，获取7日前认识请求 state说明： 0-待处理 1-已婉拒 2-过期自动拒绝 3-已同意
// @ID get_friends_out_of_seven_days
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {array} models.FriendToMeResponse
// @Router /user/relation/getFriendsOutOfSevenDays [get]
func GetFriendsOutOfSevenDays() {
	//用户可获取7日前认识请求
}

// GetMyFriendRequests
// @Summary 用户可获取对其它用户的认识请求
// @Description 可根据用户ID，获取对其它用户的认识请求 state说明： 0-待处理 1-已婉拒 2-过期自动拒绝 3-已同意
// @ID get_my_friend_requests
// @Tags Relation
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {array} models.MyFriendRequest
// @Router /user/relation/getMyFriendRequests [get]
func GetMyFriendRequests() {
	//用户可获取对其它用户的认识请求
}
