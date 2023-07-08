package injector

import (
	"go-geo-backend/src/domain/repository"
	"go-geo-backend/src/handler"
	"go-geo-backend/src/infra"
	"go-geo-backend/src/usecase"
)

func InjectDB() infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()
	return *sqlhandler
}

func InjectGeoRepository() repository.GeoRepository {
	sqlHander := InjectDB()
	return infra.NewGeoRepository(sqlHander)
}

func InjectGeoUsecase() usecase.GeoUsecase {
	GeoRepo := InjectGeoRepository()
	return usecase.NewGeoUsecase(GeoRepo)
}

func InjectGeoHandler() handler.GeoHandler {
	return handler.NewGeoHandler(InjectGeoUsecase())
}
