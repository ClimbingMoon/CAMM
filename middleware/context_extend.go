package middleware

import (
	"github.com/ClimbingMoon/camm/dto"
	"github.com/labstack/echo"
)

// ContextExtend 扩展context中间件
func ContextExtend(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &dto.CustomContext{Context: c}
		return h(cc)
	}
}
