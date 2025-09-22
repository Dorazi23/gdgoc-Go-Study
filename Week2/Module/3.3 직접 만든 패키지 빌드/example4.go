package main

import (
	"net/http"
	"strconv"

	"github.com/Week2/Module/MyPrimeNumPackage/example3"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/:number", func(c echo.Context) error {
		numStr := c.Param("number")
		n, err := strconv.Atoi(numStr)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, strconv.FormatBool(example3.IsPrime(n)))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
