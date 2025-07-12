package v2

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/BeesNestInc/CassetteOS/pkg/config"
	"fmt"
)
func (c *CasaOS) GetOptionStatus(ctx echo.Context) error {
	fmt.Println("🍀 getOptionStatus was called!")
	fmt.Println("*****")
	fmt.Println(config.AppInfo.EnableWiFiSetup)
	return ctx.JSON(http.StatusOK, config.AppInfo.EnableWiFiSetup)
}