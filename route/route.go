package route

import (
	control "github.com/ClimbingMoon/camm/control"
	middleware "github.com/ClimbingMoon/camm/middleware"
	"github.com/labstack/echo"
)

// Routes 单纯为了独立一个文件来定义route
func Routes(e *echo.Echo) error {
	// static files
	e.File("/favicon.ico", "public/img/favicon.ico")
	e.Static("/", "public")

	// 注册
	e.POST("/signup", control.Signup)
	// 登录
	e.POST("/signin", control.Signin)
	// 登出
	e.POST("/signout", control.Signout, middleware.CheckSignin)

	// user routes
	UserRoutes(e)

	return nil
}
