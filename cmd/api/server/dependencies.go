package server

import (
	propertyController "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/property"
	testController "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/testing"
	propertyService "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/service/property"
)

func resolveTestController() *testController.Controller {
	controller, err := testController.New()
	if err != nil {
		panic(err)
	}

	return controller
}

func resolvePropertyController() *propertyController.Controller {
	controller, err := propertyController.New(resolvePropertyService())
	if err != nil {
		panic(err)
	}

	return controller
}

func resolvePropertyService() *propertyService.Service {
	service, err := propertyService.New()
	if err != nil {
		panic(err)
	}

	return service
}
