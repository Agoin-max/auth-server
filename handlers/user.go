package handlers

import (
	"auth-server/conf"
	"auth-server/dbs"
	"auth-server/models"
	"auth-server/tools"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRegisterRequest struct {
	Username string `json:"user_name"` // 用户名
	Password string `json:"password"`  // 密码
}

type UserLoginRequest struct {
	Username string `json:"user_name"` // 用户名
	Password string `json:"password"`  // 密码
}

func GenerateToken(username string, user_id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user_id
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 生成签名字符串
	secret := conf.AppSetting.JwtSecret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// @Summary 用户注册
// @Schemes
// @Description 用户注册
// @Tags User
// @Param object body UserRegisterRequest false "Json请求体"
// @Success 200 {object} tools.ResponseStruct
// @Router /user/register [post]
func Register(c *gin.Context) {
	body := UserRegisterRequest{}
	if err := c.BindJSON(&body); err != nil {
		tools.Fail(c, "解析Json失败", nil)
	}

	// 校验
	if body.Username == "" || body.Password == "" {
		tools.Fail(c, "用户/密码不能为空", nil)
	}

	user := models.User{}
	user.Username = body.Username
	user.Password = tools.Md5Encode(body.Password)

	err := dbs.DB.Where("user_name=?", user.Username).First(&user).Error
	if err != gorm.ErrRecordNotFound {
		tools.Fail(c, "用户已存在", nil)
	}

	dbs.DB.Create(&user)
	msg := "注册成功"
	tools.Success(c, msg, &user)
}

// @Summary 用户登录
// @Schemes
// @Description 用户登录
// @Tags User
// @Param object body UserLoginRequest false "Json请求体"
// @Success 200 {object} tools.ResponseStruct
// @Router /user/login [post]
func Login(c *gin.Context) {
	body := UserLoginRequest{}
	if err := c.BindJSON(&body); err != nil {
		tools.Fail(c, "解析Json失败", nil)
	}

	user := models.User{}
	user.Username = body.Username
	password := tools.Md5Encode(body.Password)

	err := dbs.DB.Where("user_name=?", user.Username).First(&user).Error
	if err != gorm.ErrRecordNotFound {
		if user.Password == password {
			// 鉴权token
			token, _ := GenerateToken(body.Username, int(user.ID))
			tools.Success(c, "登录成功", token)
		} else {
			tools.Fail(c, "登录失败", nil)
		}
	} else {
		tools.Fail(c, "登录失败", nil)
	}
}

func TestAuthorization(ctx *gin.Context) {
	user := ctx.MustGet("user").(jwt.MapClaims)
	fmt.Println(user)
	tools.Success(ctx, "Auth success", user["username"])
}
