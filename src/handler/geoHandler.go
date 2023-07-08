package handler

import (
	"go-geo-backend/src/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type GeoHandler struct {
	geoUsecase usecase.GeoUsecase
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewGeoHandler(geoUsecase usecase.GeoUsecase) GeoHandler {
	geoHandler := GeoHandler{geoUsecase: geoUsecase}
	return geoHandler
}

func (handler *GeoHandler) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		postalCode := c.QueryParam("postal_code")
		if len(postalCode) != 7 {
			errorResponse := ErrorResponse{
				Message: "Invalid PostalCode",
			}
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		model, err := handler.geoUsecase.Search(postalCode)
		if err != nil {
			errorResponse := ErrorResponse{
				Message: err.Error(),
			}
			return c.JSON(http.StatusBadRequest, errorResponse)
		}
		return c.JSON(http.StatusOK, model)
	}
}

func (handler *GeoHandler) History() echo.HandlerFunc {
	return func(c echo.Context) error {
		models, err := handler.geoUsecase.History()
		if err != nil {
			errorReponse := ErrorResponse{
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, errorReponse)
		}
		return c.JSON(http.StatusOK, models)
	}
}
