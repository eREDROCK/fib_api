package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetBirthDay(c echo.Context) error {
	nstr := c.QueryParam("n")
	dateArr := strings.Split(nstr, "-")
	if len(dateArr) != 3 {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "400", "message": "Invalid request"})
	}

	year, _ := strconv.Atoi(dateArr[0])
	month, _ := strconv.Atoi(dateArr[1])
	day, _ := strconv.Atoi(dateArr[2])

	fmt.Print(year)
	fmt.Print(month)
	fmt.Print(day)

	if month == 1 || month == 2 {
		year--
		month += 12
	}

	K := year % 100
	J := year / 100

	h := (day + (13*(month+1))/5 + K + K/4 + J/4 - 2*J) % 7

	weekdays := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	return c.JSON(http.StatusOK, map[string]string{"result": weekdays[h]})

}