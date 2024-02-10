package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nairoki23/video-nairoki-backend/Database"
)

func main() {
	db := Database.GetDatabase()
	//以下API
	e := echo.New()
	e.GET("/",
		func(c echo.Context) error {
			u := &Database.Hello{
				Body: "Hello World",
			}
			return c.JSON(404, u)
		})
	e.GET("/hello", func(c echo.Context) error {
		var hellos []Database.Hello
		db.Find(&hellos)
		return c.JSON(200, hellos)
	})

	e.GET("/set-hello", func(c echo.Context) error {
		text := c.QueryParam("hello")
		hello := &Database.Hello{
			Body: text,
		}
		db.Create(&hello)
		return c.JSON(200, hello)
	})

	e.GET("/konnitiha", func(c echo.Context) error {
		var konnitihas []Database.Konnitiha
		db.Find(&konnitihas)
		return c.JSON(200, konnitihas)
	})

	e.GET("/set-konnitiha", func(c echo.Context) error {
		text := c.QueryParam("konnitiha")
		konnitiha := &Database.Konnitiha{
			Body: text,
		}
		db.Create(&konnitiha)
		return c.JSON(200, konnitiha)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
