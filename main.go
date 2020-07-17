package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	cityHandler "github.com/roysitumorang/laukpauk-marketplace/modules/city/presenter"
	provinceHandler "github.com/roysitumorang/laukpauk-marketplace/modules/province/presenter"
	subdistrictHandler "github.com/roysitumorang/laukpauk-marketplace/modules/subdistrict/presenter"
	"github.com/roysitumorang/laukpauk-marketplace/util/db"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("can't load .env file")
	}
	if err := db.Init(); err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Conn.Close()
	e := echo.New()
	if os.Getenv("DEBUG") == "1" {
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			c.Logger().Error(err)
			c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{
			echo.OPTIONS,
			echo.GET,
			echo.HEAD,
			echo.PUT,
			echo.PATCH,
			echo.POST,
			echo.DELETE,
		},
	}))

	g := e.Group("/api/v1")

	provinceGroup := g.Group("/provinces")
	provinceHandler.Mount(provinceGroup)

	cityGroup := g.Group("/cities")
	cityHandler.Mount(cityGroup)

	subdistrictGroup := g.Group("/subdistricts")
	subdistrictHandler.Mount(subdistrictGroup)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
