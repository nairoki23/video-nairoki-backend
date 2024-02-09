package main

import (
	"github.com/labstack/echo/v4"
)

type Hello struct {
	ID   uint   `gorm:"primary_key"`
	Body string `json:"body" gorm:"size:255"`
}

func main() {
	db := StartDB()
	//以下API
	e := echo.New()
	e.GET("/",
		func(c echo.Context) error {
			u := &Hello{
				Body: "Hello World",
				ID:   213,
			}
			return c.JSON(404, u)
		})
	e.GET("/hello", func(c echo.Context) error {
		var hellos []Hello
		db.Find(&hellos)
		return c.JSON(200, hellos)
	})

	e.GET("/set-hello", func(c echo.Context) error {
		text := c.QueryParam("hello")
		hello := &Hello{
			Body: text,
		}
		db.Create(&hello)
		return c.JSON(200, hello)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
