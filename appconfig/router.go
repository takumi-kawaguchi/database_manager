package appcofig

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/takumi-kawaguchi/database_manager/infrastructure"
	"github.com/takumi-kawaguchi/database_manager/interfaces/controllers"
)

// Init initializes route config
func Init() {
	e := echo.New()
	e.Renderer = infrastructure.NewTemplate()

	// router setting
	err := e.GET("/", func(c echo.Context) error {
		err2 := controllers.GetDatabases(c)
		fmt.Printf("[TEST] err2: %+v", err2)
		return err2
	})
	if err != nil {
		fmt.Println("[TEST] GetDatabases error")
		log.Printf("%+v", err)
	}

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, world")
	})

	e.Start(":1324")
}
