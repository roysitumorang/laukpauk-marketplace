package presenter

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	cityModel "github.com/roysitumorang/laukpauk-marketplace/modules/city/model"
	provinceModel "github.com/roysitumorang/laukpauk-marketplace/modules/province/model"
	"github.com/roysitumorang/laukpauk-marketplace/util/db"
)

func Mount(g *echo.Group) {
	g.GET("", getProvinces)
	g.GET("/:id/cities", getCities)
}

func getProvinces(c echo.Context) (err error) {
	var provinces []provinceModel.Province
	if err = db.Conn.Select(&provinces, `SELECT * FROM provinces ORDER BY name ASC`); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, provinces)
}

func getCities(c echo.Context) (err error) {
	var (
		provinceID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
		cities        []cityModel.City
	)
	if err = db.Conn.Select(&cities, `SELECT * FROM cities WHERE province_id = $1 ORDER BY type, name ASC`, provinceID); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, cities)
}
