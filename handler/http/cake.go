package http

import (
	"errors"
	"net/http"

	"github.com/guntoroyk/cake-store-api/entity"
	"github.com/guntoroyk/cake-store-api/lib/converter"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetCakes(c echo.Context) error {
	cakes, err := h.cakeUsecase.GetCakes()

	resp := HttpResponse{
		Code: http.StatusOK,
		Data: cakes,
	}

	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	}

	return c.JSON(resp.Code, resp)
}

func (h *handler) GetCake(c echo.Context) error {
	id := c.Param("id")

	cake, err := h.cakeUsecase.GetCake(converter.ToInt(id))

	resp := HttpResponse{}

	if errors.Is(err, entity.ErrCakeNotFound) {
		resp.Code = http.StatusNotFound
		resp.Error = err.Error()
	} else if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusOK
		resp.Data = cake
	}

	return c.JSON(resp.Code, resp)
}

func (h *handler) CreateCake(c echo.Context) error {
	cake := new(entity.Cake)
	resp := HttpResponse{}

	if err := c.Bind(cake); err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = err.Error()

		return c.JSON(resp.Code, resp)
	}

	cake, err := h.cakeUsecase.CreateCake(cake)

	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusCreated
		resp.Data = cake
	}

	return c.JSON(resp.Code, resp)
}

func (h *handler) UpdateCake(c echo.Context) error {
	cake := new(entity.Cake)
	resp := HttpResponse{}

	if err := c.Bind(cake); err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = err.Error()

		return c.JSON(resp.Code, resp)
	}

	cakeID := c.Param("id")
	cake.ID = converter.ToInt(cakeID)

	cake, err := h.cakeUsecase.UpdateCake(cake)

	if errors.Is(err, entity.ErrCakeNotFound) {
		resp.Code = http.StatusNotFound
		resp.Error = err.Error()
	} else if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusOK
		resp.Data = cake
	}

	return c.JSON(resp.Code, resp)
}

func (h *handler) DeleteCake(c echo.Context) error {
	id := c.Param("id")

	err := h.cakeUsecase.DeleteCake(converter.ToInt(id))

	resp := HttpResponse{}

	if errors.Is(err, entity.ErrCakeNotFound) {
		resp.Code = http.StatusNotFound
		resp.Error = err.Error()
	} else if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusNoContent
	}

	return c.JSON(resp.Code, resp)
}
