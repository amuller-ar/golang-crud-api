package server

import (
	testController "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/testing"
)

func resolveTestController() *testController.Controller {
	controller, err := testController.New()
	if err != nil {
		panic(err)
	}

	return controller
}
