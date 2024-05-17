package controller

import (
	"log"
	"net/http"
	"pertemuan3/model"

	"github.com/labstack/echo/v4"
)

type response struct {
	Data interface{} `"json:"data"`
	Message string `json:"message"`
}

func responseData (c echo.Context, status int, data interface{}, message string) error {
	return c.JSON(status, response{
	Data: data,
	Message: message,
	})
}

func CreateStudents(c echo.Context) (err error) {
	var s model.Students
	
	err = c.Bind(&s)
	if err != nil {
		log.Println("failed parse request")
		return
	}

	return responseData(c, http.StatusCreated, s, "Successfully Create Students")
}