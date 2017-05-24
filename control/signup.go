package control

import (
	"net/http"
	"time"

	config "github.com/ClimbingMoon/camm/config"
	m "github.com/ClimbingMoon/camm/model"
	util "github.com/ClimbingMoon/camm/util"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Signup 用户注册
func Signup(c echo.Context) error {
	if config.ConfigSource.Debug {
		log.SetLevel(log.DebugLevel)
	}

	newUser := new(m.User)
	if err := c.Bind(newUser); err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"username": newUser.Username,
		"password": newUser.Password,
		"email":    newUser.Email,
	}).Debug("User Register")

	// 用户名检查
	if !util.JudgeUsername(newUser.Username) {
		return c.String(http.StatusOK, "username check failed!")
	}
	// 邮箱检查
	if !util.JudgeEmail(newUser.Email) {
		return c.String(http.StatusOK, "email check failed!")
	}
	// TODO 错误管理???? 重复用户名或邮箱检查
	if num, _ := newUser.GetCountByUsernameOrEmail(); num > 0 {
		return c.String(http.StatusOK, "username or email conflict!")
	}
	// 密码检查
	if !util.JudgePassword(newUser.Password) {
		return c.String(http.StatusOK, "email check failed!")
	}
	// 密码计算摘要 MD5+salt
	newUser.Password = util.DigestMD5(newUser.Password + config.ConfigSource.Salt)

	// 注册时间
	newUser.CreateTime = time.Now()

	// 默认能够登录
	newUser.IsAbleLogin = true
	// 默认不是管理员 除非是第一个用户
	newUser.IsAdmin = false
	if totalUserCount, err := newUser.GetTotalCount(); totalUserCount <= 0 && err == nil {
		newUser.IsAdmin = true
	}

	newUser.Insert()
	return c.String(http.StatusOK, "OK!")
}
