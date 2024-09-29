package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api/controllers"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	// ミドルウェアを設定→本番環境用に消去
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//APIのルーティング
	e.GET("/fib", controllers.GetFib)
	e.GET("/birthday", controllers.GetBirthDay)

	return e
}
