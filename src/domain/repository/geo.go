package repository

import "go-geo-backend/src/domain/model"

type GeoRepository interface {
	GetGeo(postalCode string) (geoInfo *model.GeoInfo, err error)
	GetAccessLogs() (accessLogs []*model.AccessLog, err error)
}
