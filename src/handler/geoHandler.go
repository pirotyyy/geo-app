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

// Search API
func (handler *GeoHandler) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		postalCode := c.QueryParam("postal_code")

		model, err := handler.geoUsecase.Search(postalCode)
		if err != nil {
			errorResponse := ErrorResponse{
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, errorResponse)
		}
		return c.JSON(http.StatusOK, model)
	}
}

// AccessLog API
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
