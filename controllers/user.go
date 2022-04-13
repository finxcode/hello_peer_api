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
	CoverImageUrl string  `example:"www.coverUrl.com"`
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
	CoverImageUrl string   `example:"www.coverUrl.com"`
	Tags          string   `example:"猫控 读书达人 电影爱好者"`
	Images        []string `example:"www.imgUrl1.com, www.imgUrl2.com"`
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
	SelfDesc      string   `example:"自我描述"`
	Hobbies       string   `example:"兴趣爱好"`
	Declaration   string   `example:"交友宣言"`
	TheOne        string   `example:"希望另一半的样子"`
	Tags          string   `example:"猫控 读书达人 电影爱好者 旅行者"`
	Images        []string `example:"www.imgUrl1.com, www.imgUrl2.com"`
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
// @Success 200 {array} RandomUser
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
// @Param uid query string true "用户ID"
// @Success 200 {object} UserDetails
// @Router /user/getRandomUserDetails/{user_id} [get]
func GetRandomUserDetails(c *gin.Context) {
	//获取用户详情
}

// GetRecommendedUserList
// @Summary 获取用户推荐用户列表
// @Description 可通过用户ID获取推荐用户列表
// @ID get_recommended_user_list
// @Tags User
// @Accept application/json
// @Produce application/json
// @Security x-token
// @param x-token header string true "Authorization"
// @Success 200 {array} RecommendedUser
// @Router /user/getRecommendedUserList [get]
func GetRecommendedUserList(c *gin.Context) {
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
// @Success 200
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
// @Success 200
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
// @Success 200
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
// @Success 200
// @Router /user/upload/setUserCover [post]
func SetUserCover(c *gin.Context) {
	//设置用户头像
}
