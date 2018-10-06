package presenter

import (
	"net/http"

	"github.com/labstack/echo"
	cityModel "github.com/roysitumorang/laukpauk-marketplace/modules/city/model"
	provinceModel "github.com/roysitumorang/laukpauk-marketplace/modules/province/model"
	"github.com/roysitumorang/laukpauk-marketplace/util/db"
)

var ()

func Mount(g *echo.Group) {
	g.GET("", getProvinces)
	g.GET("/:id/cities", getCities)
}

func getProvinces(c echo.Context) error {
	provinces := []provinceModel.Province{}
	err := db.Conn.Select(&provinces, `SELECT * FROM provinces ORDER BY name ASC`)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, provinces)
}

func getCities(c echo.Context) error {
	cities := []cityModel.City{}
	err := db.Conn.Select(&cities, `SELECT * FROM cities WHERE province_id = $1 ORDER BY type, name ASC`, c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, cities)
}
