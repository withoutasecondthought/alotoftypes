package handler

import (
	"alotoftypes"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (h *Handler) getEmployee(c echo.Context) error {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := h.services.GetEmployee(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, alotoftypes.MessageJSON{Message: fmt.Sprintf("error to get employee: %s", err.Error())})
		return err
	}
	c.JSON(http.StatusOK, res)
	return nil
}

func (h *Handler) postEmployee(c echo.Context) error {
	var req alotoftypes.EmployeeJSON
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, alotoftypes.MessageJSON{Message: fmt.Sprintf("error parsing JSON: %s", err.Error())})
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := h.services.PostEmployee(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, alotoftypes.MessageJSON{Message: fmt.Sprintf("error to post employee: %s", err.Error())})
		return err
	}

	err = c.JSON(http.StatusOK, alotoftypes.MessageJSON{
		Message: "ok",
	})

	return nil
}

func (h *Handler) putEmployee(c echo.Context) error {
	var req alotoftypes.EmployeeUpdateJSON

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, alotoftypes.MessageJSON{Message: fmt.Sprintf("error reading BODY: %s", err.Error())})
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := h.services.PutEmployee(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, alotoftypes.MessageJSON{Message: fmt.Sprintf("error to put employee: %s", err.Error())})
		return err
	}

	err = c.JSON(http.StatusOK, alotoftypes.MessageJSON{
		Message: "ok",
	})

	err = c.JSON(http.StatusOK, req)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) deleteEmployee(c echo.Context) error {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := h.services.DeleteEmployee(ctx, id)
	if err != nil {
		err := c.JSON(http.StatusBadRequest, alotoftypes.MessageJSON{Message: fmt.Sprintf("error to get employee: %s", err.Error())})
		return err
	}
	err = c.JSON(http.StatusOK, alotoftypes.MessageJSON{
		Message: "ok",
	})
	if err != nil {
		return err
	}
	return nil
}
