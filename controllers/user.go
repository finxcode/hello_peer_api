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
	CoverImageUrl string  `example:"www.baidu.com"`
}

type RandomUserDetails struct {
	UserName      string   `example:"豆豆"`
	Age           int      `example:"25"`
	Occupation    string   `example:"平面设计师"`
	Constellation string   `example:"处女座"`
	Height        string   `example:"165cm"`
	Weight        string   `example:"43kg"`
	Education     string   `example:"本科"`
	Resident      string   `example:"深圳"`
	Hometown      string   `example:"湖南长沙"`
	SelfDesc      string   `example:"自我描述"`
	Hobbies       string   `example:"兴趣爱好"`
	Declaration   string   `example:"交友宣言"`
	TheOne        string   `example:"希望另一半的样子"`
	Tags          string   `example:"猫控 读书达人 电影爱好者 旅行者"`
	Images        []string `example:"www.baidu.com, www.bing.com"`
}

// GetRandomUsersInSquare
// @Summary 获取用户广场随机用户列表
// @Description 可通过用户ID以及用户广场设置项获取随机用户列表
// @ID get_random_users_list_in_square
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param uid query string true "用户ID"
// @Param pagination body models.Pagination true "分页"
// @Success 200 {array} RandomUser
// @Router /user/getRandomUsers/{user_id} [post]
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
// @Success 200 {object} RandomUserDetails
// @Router /user/getRandomUsers/{user_id} [get]
func GetRandomUserDetails(c *gin.Context) {
	//获取用户详情
}

// GetSquareSettings
// @Summary 获取用户广场设置
// @Description 可通过用户ID获取用户广场设置
// @ID get_square_settings
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param uid query string true "用户ID"
// @Success 200 {object} models.SquareSettings
// @Router /user/getSquareSettings/{user_id} [get]
func GetSquareSettings(c *gin.Context) {
	//获取用户广场设置
}

// SetSquareSettings
// @Summary 设置用户广场设置
// @Description 可设置用户广场设置
// @ID set_square_settings
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param uid query string true "用户ID"
// @Param settings body models.SquareSettings true "设置"
// @Success 200
// @Router /user/setSquareSettings/{user_id} [post]
func SetSquareSettings(c *gin.Context) {
	//设置用户广场设置
}
