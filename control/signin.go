package control

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	config "github.com/ClimbingMoon/camm/config"
	m "github.com/ClimbingMoon/camm/model"
	util "github.com/ClimbingMoon/camm/util"
	session "github.com/ipfans/echo-session"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Signin 用户登录
func Signin(c echo.Context) error {
	session := session.Default(c)

	if config.ConfigSource.Debug {
		log.SetLevel(log.DebugLevel)
	}

	sessionUserID := session.Get("id")
	if sessionUserID != nil {
		log.WithField("session_user_id", sessionUserID).Debug("id in session")
		sessionUser := m.User{ID: bson.ObjectIdHex(sessionUserID.(string))}
		if alreadyUser, err := (sessionUser.GetOneByID()); err == nil {
			log.WithField("id", alreadyUser.ID).Debug("user already logged in")
			return c.String(http.StatusOK, "user already logged in")
		}
	}

	loginUser := new(m.User)
	if err := c.Bind(loginUser); err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"username": loginUser.Username,
		"password": loginUser.Password,
	}).Debug("User Login")

	// 用户名检查
	if !util.JudgeUsername(loginUser.Username) {
		return c.String(http.StatusOK, "username check failed!")
	}
	// 密码检查
	if !util.JudgePassword(loginUser.Password) {
		return c.String(http.StatusOK, "email check failed!")
	}
	// 检查用户是否存在
	var targetUser m.User
	var err error
	if targetUser, err = loginUser.GetOneByUsername(); err != nil {
		log.WithField("username", loginUser.Username).Error("user " + err.Error())
		return c.String(http.StatusOK, "user "+err.Error())
	}

	log.WithFields(log.Fields{
		"id":       targetUser.ID,
		"username": targetUser.Username,
	}).Debug("finding user")

	// 密码计算摘要 MD5+salt
	inputPassword := util.DigestMD5(loginUser.Password + config.ConfigSource.Salt)
	// 检查密码是否正确
	if targetUser.Password != inputPassword {
		return c.String(http.StatusOK, "wrong password")
	}

	// 密码正确 存储登录用户id
	// TODO 删除已有session 这需要 token持久机制 或 对redis的操作
	session.Set("id", targetUser.ID.Hex())
	session.Save()

	return c.String(http.StatusOK, "Signin OK!")
}
