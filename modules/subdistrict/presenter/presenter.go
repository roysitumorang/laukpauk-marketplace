package presenter

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	villageModel "github.com/roysitumorang/laukpauk-marketplace/modules/village/model"
	"github.com/roysitumorang/laukpauk-marketplace/util/db"
)

func Mount(g *echo.Group) {
	g.GET("/:id/villages", getVillages)
}

func getVillages(c echo.Context) (err error) {
	var (
		subdistrictID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
		villages         []villageModel.Village
	)
	if err = db.Conn.Select(&villages, `SELECT * FROM villages WHERE subdistrict_id = $1 ORDER BY name ASC`, subdistrictID); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, villages)
}
