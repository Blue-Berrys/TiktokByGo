package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/middleware/JWT"
	"github.com/neverTanking/TiktokByGo/model"
	"net/http"
	"strconv"
)

type User struct {
	model.User
	IsFollow bool `json:"is_follow"` // true-已关注,false-未关注
}

type UserLoginResponse struct {
	Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User model.User `json:"user"`
}

// 验证密码
func verifyPassword(input string, password string) bool {

	return input == password
}

func isUsernameValid(username string) bool {
	return len(username) > 0 && len(username) <= 32
}
func isPasswordValid(password string) bool {
	return len(password) > 0 && len(password) <= 32
}

// Register 用户登录处理函数，返回用户登录信息
func Register(c *gin.Context) {
	// 1. http请求中获取用户信息验证合法性
	username := c.Query("username")
	password := c.Query("password")
	if !(isUsernameValid(username) && isPasswordValid(password)) {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: Fail, StatusMsg: Valid},
		})
	}

	// 2. 查询用户是否存在，并返回用户信息
	_, exist := model.SearchUserByName(username)
	// 3.1 存在则报错
	if exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: Fail, StatusMsg: Existed},
		})
	}
	// 3.2 不存在则创建用户，并返回用户信息
	userId, err := model.CreatUser(username, password)
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: Fail, StatusMsg: UnknownReason},
		})
	}
	// 4.生成token
	token, err := JWT.GetToken(userId, username, password)
	if err != nil {
		fmt.Println(fmt.Errorf("token generate err: %v", err))
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: Fail, StatusMsg: TokenFail},
		})
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: Success, StatusMsg: SignUpOk},
		UserId:   userId,
		Token:    token,
	})
}

// Login 用户登录处理函数，返回用户登录信息
func Login(c *gin.Context) {
	// 1. http请求中获取用户信息并检查输入合法性
	username := c.Query("username")
	password := c.Query("password")
	if !(isUsernameValid(username) && isPasswordValid(password)) {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: Fail, StatusMsg: Valid},
		})
	}

	// 2. 查询用户是否存在，并返回用户信息
	user, exist := model.SearchUserByName(username)
	if !exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: Fail, StatusMsg: NotExisted},
		})
		return
	}
	//3. 验证密码
	if !verifyPassword(password, user.Password) {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: Fail, StatusMsg: WrongPassword},
		})
		return
	}

	// 4. 生成token
	token, err := JWT.GetToken(user.ID, username, password)
	if err != nil {
		fmt.Print(fmt.Errorf("token generate fail: %v", err))
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: Fail, StatusMsg: TokenFail},
		})
	}

	// 5. 返回用户信息和token
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: Success, StatusMsg: LoginOk},
		UserId:   user.ID,
		Token:    token,
	})
}

// UserInfo description: 获取用户信息处理函数，返回用户信息
func UserInfo(c *gin.Context) {

	var userId, _ = strconv.ParseUint(c.Query("user_id"), 10, 64)

	user, ok := model.SearchUserByID(uint(userId))
	if ok != true {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: Fail, StatusMsg: NotExisted},
		})
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: Success},
		User:     user,
	})
}
