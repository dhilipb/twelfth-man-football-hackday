package handlers

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func Chants(c *echo.Context) {
	team := c.Param("team")

	var f *os.File
	var err error
	if f, err = os.Open("fixtures/" + team + ".json"); err != nil {
		c.Response.WriteHeader(http.StatusNotFound)
		return
	}

	chants := make([]byte, 20*1024)
	if n, err := f.Read(chants); err != nil {
		c.Response.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		chants = chants[0:n]
	}

	c.Response.WriteHeader(http.StatusOK)
	c.Response.Write(chants)
}
