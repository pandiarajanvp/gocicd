package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type Person struct {
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

func getPersonDetails(c echo.Context) error {
	p := Person{}
	p.Avatar = "https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/50.jpg"
	p.CreatedAt = "2025-03-24T14:50:23.348Z"
	p.Id = "1"
	p.Name = "Pandian"

	return c.JSON(http.StatusOK, p)
}
