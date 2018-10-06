package presenter

import (
	"net/http"

	"github.com/labstack/echo"
	villageModel "github.com/roysitumorang/laukpauk-marketplace/modules/village/model"
	"github.com/roysitumorang/laukpauk-marketplace/util/db"
)

var ()

func Mount(g *echo.Group) {
	g.GET("/:id/villages", getVillages)
}

func getVillages(c echo.Context) error {
	villages := []villageModel.Village{}
	err := db.Conn.Select(&villages, `SELECT * FROM villages WHERE subdistrict_id = $1 ORDER BY name ASC`, c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, villages)
}
