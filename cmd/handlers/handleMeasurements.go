package handlers

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func HandleCreateMeasurement(c echo.Context) error {
	measurement := models.Measurements{}

	c.Bind(&measurement)

	newMeasurement, err := repositories.CreateMeasurement(measurement)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusAccepted, newMeasurement)
}

func HandleUpdateUser(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user := models.User{}
	c.Bind(&user)

	updatedUser, err := repositories.UpdateUser(user, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}
