package server

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/auth"
	propertyController "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/property"
	userController "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/user"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/database"
	propertyRepository "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/repository/property"
	userRepository "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/repository/user"
	propertyService "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/service/property"
	userService "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/service/user"
	"gorm.io/gorm"
)

func resolvePropertyController() *propertyController.Controller {
	controller, err := propertyController.New(resolvePropertyService())
	checkErr(err)

	return controller
}

func resolveUserController() *userController.Controller {
	controller, err := userController.New(
		resolveUserService(),
		auth.NewTokenService(),
	)
	checkErr(err)

	return controller
}

func resolvePropertyService() *propertyService.Service {
	service, err := propertyService.New(resolvePropertyRepository())
	checkErr(err)

	return service
}

func resolveUserService() *userService.Service {
	service, err := userService.New(
		resolveUserRepository(),
		resolvePropertyRepository(),
	)
	checkErr(err)

	return service
}

func resolvePropertyRepository() *propertyRepository.Repository {
	repo, err := propertyRepository.New(resolveSqlClient())
	checkErr(err)

	return repo
}

func resolveUserRepository() *userRepository.Repository {
	repo, err := userRepository.New(resolveSqlClient())
	checkErr(err)

	return repo
}

func resolveSqlClient() *gorm.DB {
	return database.GetDB()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
