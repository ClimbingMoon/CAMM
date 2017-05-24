package dto

import (
	"github.com/ClimbingMoon/camm/model"
	"github.com/labstack/echo"
)

// CustomContext 扩展echo上下文
type CustomContext struct {
	echo.Context // echo默认

	User *model.User // 用户信息
}
