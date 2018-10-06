package presenter

import (
	"net/http"

	"github.com/labstack/echo"
	subdistrictModel "github.com/roysitumorang/laukpauk-marketplace/modules/subdistrict/model"
	"github.com/roysitumorang/laukpauk-marketplace/util/db"
)

var ()

func Mount(g *echo.Group) {
	g.GET("/:id/subdistricts", getSubdistricts)
}

func getSubdistricts(c echo.Context) error {
	subdistricts := []subdistrictModel.Subdistrict{}
	err := db.Conn.Select(&subdistricts, `SELECT * FROM subdistricts WHERE city_id = $1 ORDER BY name ASC`, c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, subdistricts)
}
