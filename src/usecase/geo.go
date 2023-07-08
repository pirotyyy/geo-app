package usecase

import (
	"go-geo-backend/src/domain/model"
	"go-geo-backend/src/domain/repository"
)

// Define Usecase
type GeoUsecase interface {
	Search(string) (geoInfo *model.GeoInfo, err error)
	History() (accessLogs []*model.AccessLog, err error)
}

type geoUsecase struct {
	geoRepo repository.GeoRepository
}

func NewGeoUsecase(geoRepo repository.GeoRepository) GeoUsecase {
	geoUsecase := geoUsecase{geoRepo: geoRepo}
	return &geoUsecase
}

func (usecase *geoUsecase) Search(postalCode string) (geoInfo *model.GeoInfo, err error) {
	geoInfo, err = usecase.geoRepo.GetGeo(postalCode)
	return
}

func (usecase *geoUsecase) History() (accessLogs []*model.AccessLog, err error) {
	accessLogs, err = usecase.geoRepo.GetAccessLogs()
	return
}
