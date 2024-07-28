package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// フィボナッチ数列を計算する関数
func solve_fibonacci(n int) int {
	var Fn1, Fn0 int
	if n == 0 || n == 1 {
		return n
	} else if n > 1 {
		Fn0 = 0
		Fn1 = 1
		for i := 0; i < n/2; i++ {
			Fn0 = Fn0 + Fn1
			Fn1 = Fn0 + Fn1
		}
	}

	if n%2 == 0 {
		return Fn0
	} else {
		return Fn1
	}
}

// フィボナッチ数列を計算して返すハンドラー
func GetFib(c echo.Context) error {
	nStr := c.QueryParam("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n < 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "400", "message": "Invalid request"})
	}

	// n番目のフィボナッチ数列を計算
	result := solve_fibonacci(n)
	return c.JSON(http.StatusOK, map[string]int{"result": result})
}
