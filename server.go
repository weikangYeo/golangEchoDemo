package main
import (
	"github.com/labstack/echo"
	"repo/test_go_echo/controller/staff"
	"github.com/labstack/echo/middleware"
	"net/http"
	"github.com/labstack/gommon/log"
	//"os"
	"github.com/dgrijalva/jwt-go"
	"time"
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

	//write to file
	//file, _ := os.Create("log.txt")
	//e.Logger.SetOutput(file)

	//gommon log, header will diff from e.Logger
	log.SetHeader("[GOMMON LOG] ${time_rfc3339} ${level}")
	log.Info("test here")
	e.Use(middleware.Recover())
	log.Debug("Created user")


	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	//}))

	//e.GET("/", func(c echo.Context) error {
	//	c.Logger().Info("test when calling API")
	//	return c.JSON(http.StatusCreated, "Welcome mvc echo with mysql app using Golang")
	//})

	//test authentication with JWT
	e.GET("/employees", staff.GetEmployees2)
	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group,
	// mean the following group will run JWT
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)
	r.GET("/test", func(c echo.Context) error {
			return c.JSON(http.StatusCreated, "Welcome mvc echo with mysql app using Golang with jwt auth")
		})
	e.Logger.Fatal(e.Start(":8081"))



}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	c.Logger().Info("username: " )
	c.Logger().Info(username )

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Minute).Unix()
	//claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}