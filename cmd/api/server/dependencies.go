package server

import (
	propertyController "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/property"
	testController "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/testing"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/repository/property"
	propertyService "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/service/property"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func resolveTestController() *testController.Controller {
	controller, err := testController.New()
	checkErr(err)

	return controller
}

func resolvePropertyController() *propertyController.Controller {
	controller, err := propertyController.New(resolvePropertyService())
	checkErr(err)

	return controller
}

func resolvePropertyService() *propertyService.Service {
	service, err := propertyService.New(resolvePropertyRepository())
	checkErr(err)

	return service
}

func resolvePropertyRepository() *property.Repository {
	repo, err := property.New(resolveSqlClient())
	checkErr(err)

	return repo
}

func resolveSqlClient() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
