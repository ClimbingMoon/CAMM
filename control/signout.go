package control

import (
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

// Signout 用户登出
func Signout(c echo.Context) error {

	// 清除session
	session := session.Default(c)
	session.Clear()
	session.Save()

	return c.String(http.StatusOK, "Successfully sign out!")
}
