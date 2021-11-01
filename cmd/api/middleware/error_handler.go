package middleware

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) == 0 {
		return
	}

	var restError rest.Error
	var errors []error
	var hasRestError bool

	for _, e := range ctx.Errors {
		if restErr, ok := e.Err.(rest.Error); ok {
			restError = restErr
			hasRestError = ok

			break
		}

		errors = append(errors, e)
	}

	if !hasRestError {
		restError = rest.NewError(
			http.StatusInternalServerError,
			ctx.Err().Error(),
			errors...,
		)
	}

	ctx.JSON(restError.Status, restError)

	if !ctx.IsAborted() {
		ctx.AbortWithStatus(restError.Status)
	}
}
