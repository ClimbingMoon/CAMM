package main

import (
	config "github.com/ClimbingMoon/camm/config"
	middleware "github.com/ClimbingMoon/camm/middleware"
	route "github.com/ClimbingMoon/camm/route"
	session "github.com/carynova/echo-session"

	echomw "github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func main() {
	configSource := config.ConfigSource

	e := echo.New()

	// 注册middleware
	e.Use(echomw.Gzip())
	e.Use(middleware.Logrus())
	if !configSource.Debug {
		e.Use(echomw.Recover())
	}

	// redis session
	store, err := session.NewRedisStore(32, "tcp", configSource.RedisHost+":"+configSource.RedisPort, "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	store.Options(session.Options{
		MaxAge: 7200, // 设定session缓存2小时失效
	})
	e.Use(session.Sessions("GSESSION", store))

	// 扩展context
	e.Use(middleware.ContextExtend)

	// apply route
	route.Routes(e)

	e.Logger.Fatal(e.Start(configSource.Host + ":" + configSource.Port))
}
