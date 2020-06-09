package staff

import (
	"repo/test_go_echo/models"
	"net/http"
	"github.com/labstack/echo"
)

func GetEmployees(c echo.Context) error {
	result := models.GetEmployee()
	name := c.QueryParam("name")
	println("foo")
	println("name")
	println(name)
	return c.JSON(http.StatusOK,result)
}
