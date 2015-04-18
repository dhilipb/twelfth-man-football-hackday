package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ifraixedes/twelveman/apihack"
	"github.com/labstack/echo"
)

func Matches(c *echo.Context) {
	res, err := apihack.Matches(c.Param("league"))
	if err != nil {
		errMsg := fmt.Sprintf("API HackDay returnned: %s", err.Error())
		http.Error(c.Response, errMsg, http.StatusInternalServerError)
		return
	}

	if res.StatusCode > 299 {
		errBody := NewAppErrResBody(fmt.Sprintf("API HackDay returned non statisfactory status code %d", res.StatusCode))
		c.JSON(res.StatusCode, errBody)
		return
	}

	if err != nil {
		errBody := NewAppErrResBody(fmt.Sprintf("Error reading response body from API Hackday %d", err.Error()))
		c.JSON(http.StatusInternalServerError, errBody)
		return
	}

	c.Response.WriteHeader(http.StatusOK)
	io.Copy(c.Response, res.Body)
}
