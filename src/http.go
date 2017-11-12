package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Routing
	e.POST("/edges", PostHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

type Edge struct {
	Source string  `json:"source"`
	Dest   string  `json:"dest"`
	Weight float64 `json:"weight"`
}

type Edges struct {
	Data []Edge `json:"data"`
}

func PostHandler(c echo.Context) error {
	post := new(Edges)
	if err := c.Bind(post); err != nil {
		return err
	}
	fmt.Print(post)
	return c.JSON(http.StatusCreated, post)
}
