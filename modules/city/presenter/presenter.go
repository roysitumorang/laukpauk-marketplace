package presenter

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	subdistrictModel "github.com/roysitumorang/laukpauk-marketplace/modules/subdistrict/model"
	"github.com/roysitumorang/laukpauk-marketplace/util/db"
)

func Mount(g *echo.Group) {
	g.GET("/:id/subdistricts", getSubdistricts)
}

func getSubdistricts(c echo.Context) (err error) {
	var (
		cityID, _    = strconv.ParseInt(c.Param("id"), 10, 64)
		subdistricts []subdistrictModel.Subdistrict
	)
	if err = db.Conn.Select(&subdistricts, `SELECT * FROM subdistricts WHERE city_id = $1 ORDER BY name ASC`, cityID); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, subdistricts)
}
