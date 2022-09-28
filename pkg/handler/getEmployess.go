package handler

import (
	"alotoftypes"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) getEmployeesWithParams(c echo.Context) error {
	p := c.Param("page")
	if p == "" {
		p = "0"
	}
	page, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, alotoftypes.MessageJSON{Message: fmt.Sprintf("You entered string not number of page: %s", err.Error())})
		return err
	}
	var params interface{}
	err = c.Bind(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, alotoftypes.MessageJSON{Message: fmt.Sprintf("Error parsingJSON: %s", err.Error())})
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := h.services.GetEmployeeWithParams(ctx, uint(page), params)
	if err != nil {
		c.JSON(http.StatusBadRequest, alotoftypes.MessageJSON{Message: fmt.Sprintf("request FICKING ERROR, check your mind: %s", err.Error())})
		return err
	}

	c.JSON(http.StatusOK, res)
	if err != nil {
		return err
	}
	return nil
}
