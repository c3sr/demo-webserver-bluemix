//go:generate esc -o assets.go -pkg web -private ./assets

package web

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func Start(addr string) {
	e := echo.New()

	// Setting up the termination timeout to 30 seconds.
	e.ShutdownTimeout = 30 * time.Second

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, DISCVR!")
	})
	e.GET("/version", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"version_info": versionInfo,
			"build_time":   buildTimeInfo,
			"repository":   repository,
		})
	})
	e.GET("/assets/*",
		echo.WrapHandler(http.FileServer(_escFS(false))))
	e.Logger.Fatal(e.Start(addr))
}
