package route

import (
	control "github.com/ClimbingMoon/camm/control"
	middleware "github.com/ClimbingMoon/camm/middleware"
	"github.com/labstack/echo"
)

// UserRoutes user route
func UserRoutes(e *echo.Echo) error {
	// static files
	g := e.Group("/user")
	g.Use(middleware.CheckSignin)

	// GET /user/profile 用户获取自己的信息
	g.GET("/profile", control.GetUserProfile)

	// GET /user/:id 用户获取他人信息
	g.GET("/:id", control.GetUserProfileByID)

	return nil
}
