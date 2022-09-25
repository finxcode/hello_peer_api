package controllers

import "github.com/gin-gonic/gin"

type RandomUser struct {
	UserName      string  `example:"豆豆"`
	PetName       string  `example:"Amy"`
	Age           int     `example:"25"`
	Occupation    string  `example:"平面设计师"`
	Lng           float32 `example:"113.95"`
	Lat           float32 `example:"22.51"`
	Location      string  `example:"南山区"`
	CoverImageUrl string  `example:"cover.png"`
}

type RecommendedUser struct {
	UserName      string   `example:"豆豆"`
	PetName       string   `example:"Amy"`
	Age           int      `example:"25"`
	Occupation    string   `example:"平面设计师"`
	Lng           float32  `example:"113.95"`
	Lat           float32  `example:"22.51"`
	Location      string   `example:"南山区"`
	Verified      bool     `example:"true"`
	CoverImageUrl string   `example:"cover.png"`
	Tags          string   `example:"猫控 读书达人 电影爱好者"`
	Images        []string `example:"img1.jpg,img2.jpg"`
}

type UserDetails struct {
	UserName      string   `example:"豆豆"`
	Age           int      `example:"25"`
	Occupation    string   `example:"平面设计师"`
	Constellation string   `example:"处女座"`
	Height        string   `example:"165cm"`
	Weight        string   `example:"43kg"`
	Education     string   `example:"本科"`
	Location      string   `example:"深圳"`
	Hometown      string   `example:"湖南长沙"`
	Self_Desc     string   `example:"自我描述"`
	Hobbies       string   `example:"兴趣爱好"`
	Declaration   string   `example:"交友宣言"`
	The_One       string   `example:"希望另一半的样子"`
	Tags          []string `example:"猫控,读书达人,电影爱好者,旅行者"`
	Images        []string `example:"img1.jpg,img2.jpg"`
	Cover_Image   string   `example:"cover.png"`
	Birthday      string   `json:"birthday" example:"1988-10-2"`
	Purpose       string   `json:"purpose" example:"交友"`
	Gender        int      `json:"gender" example:"1"`
	Marriage      string   `json:"marriage" example:"未婚"`
	Income        string   `json:"income" example:"15-25w"`
	PetName       string   `json:"petName" examle:"一个宠物"`
	FocusStatus   int      `json:"focusStatus" example:"1"`
	FriendStatus  int      `json:"friendStatus" example:"0"`
}

type UserDetailsUpdate struct {
	UserName      string `example:"豆豆"`
	Age           int    `example:"25"`
	Occupation    string `example:"平面设计师"`
	Constellation string `example:"处女座"`
	Height        string `example:"165cm"`
	Weight        string `example:"43kg"`
	Education     string `example:"本科"`
	Location      string `example:"深圳"`
	Hometown      string `example:"湖南长沙"`
	Self_Desc     string `example:"自我描述"`
	Hobbies       string `example:"兴趣爱好"`
	Declaration   string `example:"交友宣言"`
	The_One       string `example:"希望另一半的样子"`
	Tags          string `example:"猫控 读书达人 电影爱好者 旅行者"`
	Birthday      string `json:"birthday" example:"1988-10-2"`
	Purpose       string `json:"purpose" example:"交友"`
	Gender        int    `json:"gender" example:"1"`
	Marriage      string `json:"marriage" example:"未婚"`
	Income        string `json:"income" example:"15-25w"`
}

// GetRandomUsersInSquare
// @Summary 获取用户广场随机用户列表
// @Description 可通过用户ID以及用户广场设置项获取随机用户列表
// @ID get_random_users_list_in_square
// @Tags User
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Param pagination body models.Pagination true "分页"
// @Success 0 {array} RandomUser
// @Router /user/getRandomUsers [post]
func GetRandomUsersInSquare(c *gin.Context) {
	//获取广场随机用户接口
}

// GetRandomUserDetails
// @Summary 获取广场用户详情
// @Description 可通过用户ID获取用户详情
// @ID get_random_users_details
// @Tags User
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Param uid query string true "用户ID"
// @Success 0 {object} UserDetails
// @Router /user/getRandomUserDetails [get]
func GetRandomUserDetails(c *gin.Context) {
	//获取用户详情
}

// GetRecommendedUsers
// @Summary 获取用户推荐用户列表
// @Description 可通过用户ID获取推荐用户列表
// @ID get_recommended_users
// @Tags User
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {array} RecommendedUser
// @Router /user/getRecommendedUsers [get]
func GetRecommendedUsers(c *gin.Context) {
	//获取某一用户推荐用户列表
}

//SetUserGender
// @Summary 设置用户性别
// @Description 可通过用户token设置用户性别
// @ID set_user_gender
// @Tags User
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Param gender body models.UserGender true "用户性别"
// @Success 0
// @Router /user/setUserGender [post]
func SetUserGender(c *gin.Context) {
}

//SetUserBasicInfo
// @Summary 设置用户基础信息
// @Description 可通过用户token设置用户基础信息
// @ID set_user_basic_info
// @Tags User
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Param basicInfo body models.BasicInfo true "用户基础信息"
// @Success 0
// @Router /user/setUserBasicInfo [post]
func SetUserBasicInfo(c *gin.Context) {

}

// SetUserAvatar
// @Summary 设置用户头像
// @Description 用户可设置头像
// @ID set_user_avatar
// @Tags User
// @Accept mpfd
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @param content formData file true "头像文件"
// @Success 0
// @Router /user/upload/setUserAvatar [post]
func SetUserAvatar(c *gin.Context) {
	//设置用户头像
}

// SetUserCover
// @Summary 设置用户封面图
// @Description 用户可设置封面图
// @ID set_user_cover
// @Tags User
// @Accept mpfd
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @param content formData file true "封面文件"
// @Success 0
// @Router /user/upload/setUserCover [post]
func SetUserCover(c *gin.Context) {
	//设置用户头像
}

// SetUserImage
// @Summary 设置用户图片
// @Description 用户可上传图片
// @ID set_user_image
// @Tags User
// @Accept mpfd
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @param content formData file true "用户图片"
// @Success 0
// @Router /user/upload/setUserImage [post]
func SetUserImage(c *gin.Context) {
	//设置用户头像
}

// GetUserDetails
// @Summary 获取用户详情
// @Description 可获取用户详情
// @ID get_user_details
// @Tags User
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {object} UserDetails
// @Router /user/getUserDetails [get]
func GetUserDetails(c *gin.Context) {
	//获取用户详情
}

// GetImage
// @Summary 获取图片
// @Description 获取图片
// @ID get_image
// @Tags Util
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0
// @Router /images/{file_name} [get]
func GetImage(c *gin.Context) {
	// 获取图片
}

// SetUserDetails
// @Summary 设置用户信息
// @Description 用户可设置信息
// @ID set_user_details
// @Tags User
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Param user body UserDetailsUpdate true "用户信息"
// @Success 0
// @Router /user/setUserDetails [post]
func SetUserDetails(c *gin.Context) {
	// 更新用户信息
}

// DeleteUserImage
// @Summary 删除用户图片
// @Description 删除用户图片
// @ID delete_user_image
// @Tags User
// @Security x-token
// @param x-token header string true "Authorization"
// @param filename query string true "文件名"
// @Success 0
// @Router /user/deleteUserImage [post]
func DeleteUserImage(c *gin.Context) {
	// 删除用户图片
}

// GetUserHomepageInfo
// @Summary 获取用户首页信息
// @Description 可获取用户首页信息
// @ID get_user_homepage_info
// @Tags User
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {object} models.HomepageInfo
// @Router /user/getUserHomepageInfo [get]
func GetUserHomepageInfo(c *gin.Context) {
	// 获取用户首页信息
}

// GetUserDetailsById
// @Summary 根据用户ID获取用户详情
// @Description 可根据用户ID获取用户详情 focusStatus说明： 0-未关注 1-已关注 friendStatus说明： -1-想认识ta 0-已申请 1-发消息 2-去同意
// @ID get_user_details_by_id
// @Tags User
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @param uid query string true "user_id"
// @Success 0 {object} UserDetails
// @Router /user/getUserDetails [get]
func GetUserDetailsById(c *gin.Context) {
	//根据用户ID获取用户详情
}

// GetUserInfoCompleteLevel
// @Summary 根据用户ID获取用户信息完成度
// @Description 可根据用户ID获取用户信息完成度 0-未完成 1-基本信息 2-基本信息+头像
// @ID get_user_info_complete_level
// @Tags User
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 0 {object} models.InfoCompletionLevel
// @Router /user/getUserInfoCompleteLevel [get]
func GetUserInfoCompleteLevel(c *gin.Context) {
	//根据用户ID获取用户详情
}
