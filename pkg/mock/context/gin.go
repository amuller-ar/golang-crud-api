package context

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MakeGinContext(t *testing.T) *gin.Context {
	t.Helper()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	return ctx
}

func AddBodyRequest(ctx *gin.Context, method string, params gin.Params, body io.Reader) {
	ctx.Request, _ = http.NewRequest(method, "/", body)
	ctx.Params = params
	ctx.Request.URL.RawQuery = ctx.Request.URL.Query().Encode()
}
