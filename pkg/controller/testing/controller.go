package test

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func New() (*Controller, error) {
	return &Controller{}, nil
}

func (ctrl *Controller) GET(ctx *gin.Context) error {

	log.Println("[test]: test log")
	ctx.JSON(http.StatusOK, "esto es una prueba")
	return nil
}
