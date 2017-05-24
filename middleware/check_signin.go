package middleware

import (
	"net/http"

	session "github.com/carynova/echo-session"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	"github.com/ClimbingMoon/camm/dto"
	m "github.com/ClimbingMoon/camm/model"
	log "github.com/sirupsen/logrus"
)

// CheckSignin 检查用户是否登录的middleware
func CheckSignin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// session 取出 id
		session := session.Default(c)

		sessionUserID := session.Get("id")
		if sessionUserID != nil {
			log.WithField("session_user_id", sessionUserID).Debug("id in session")
			sessionUser := m.User{ID: bson.ObjectIdHex(sessionUserID.(string))}
			if alreadyUser, err := (sessionUser.GetOneByID()); err == nil {
				log.WithField("id", alreadyUser.ID).Debug("user already logged in")
				// user信息写入context
				c.(*dto.CustomContext).User = &alreadyUser
				return next(c)
			}
		}

		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

}
