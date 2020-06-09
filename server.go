package main
import (
	"github.com/labstack/echo"
	"repo/test_go_echo/controller/staff"
	"github.com/labstack/echo/middleware"
	"net/http"
"github.com/labstack/gommon/log"

)

func main() {
	e := echo.New()
	//e.Use(middleware.Logger())

	//middleware logging, log when api log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[Middleware LOG] time:${time_rfc3339_nano} method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("[e.Logger] ${time_rfc3339} ${level} ---> ")
	e.Logger.Info("test")

	//gommon log, header will diff from e.Logger
	log.SetHeader("[GOMMON LOG] ${time_rfc3339} ${level}")
	log.Info("test here")
	e.Use(middleware.Recover())
	log.Debug("Created user")


	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	//}))

	e.GET("/", func(c echo.Context) error {
		c.Logger().Info("test when calling API")
		return c.JSON(http.StatusCreated, "Welcome mvc echo with mysql app using Golang")
	})

	e.GET("/employees", staff.GetEmployees2)

	e.Logger.Fatal(e.Start(":8081"))

	//e.Logger.Fatal(e.Start(":1323"))


}
