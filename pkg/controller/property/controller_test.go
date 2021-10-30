package property

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/rest"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/mock"
	mockContext "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/mock/context"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/mock/domain/factory"
	mockService "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/mock/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
)

func TestController_Create(t *testing.T) {
	c := mockContext.MakeGinContext(t)

	type testCase struct {
		name      string
		configReq func(ctx *gin.Context)
		mocks     func() *mockService.PropertyService
		asserts   func(interface{})
	}

	tests := []testCase{
		{
			name:      "bad request",
			configReq: func(ctx *gin.Context) {},
			mocks: func() *mockService.PropertyService {
				return &mockService.PropertyService{}
			},
			asserts: func(got interface{}) {
				mockContext.AddBodyRequest(c, http.MethodPost, nil, nil)

				assert.Equal(t, http.StatusBadRequest, got.(rest.Error).Status)
				assert.Equal(t, "error binding json request. cause: invalid request", got.(rest.Error).Message)
			},
		},
		{
			name: "internal server error",
			configReq: func(ctx *gin.Context) {
				requestJSON, _ := mock.GetResource("create_property_request.json", mock.ResourceTypeRequest)

				mockContext.AddBodyRequest(ctx, http.MethodPost, nil, bytes.NewReader(requestJSON))
			},
			mocks: func() *mockService.PropertyService {
				propService := &mockService.PropertyService{}
				propService.
					On("Create", testifyMock.Anything).
					Return(
						nil, errors.New("some error"),
					)

				return propService
			},
			asserts: func(got interface{}) {
				assert.Equal(t, rest.NewError(http.StatusInternalServerError, "some error"), got)
			},
		},
		{
			name: "success",
			configReq: func(ctx *gin.Context) {
				requestJSON, _ := mock.GetResource("create_property_request.json", mock.ResourceTypeRequest)

				mockContext.AddBodyRequest(ctx, http.MethodPost, nil, bytes.NewReader(requestJSON))
			},
			mocks: func() *mockService.PropertyService {
				propService := &mockService.PropertyService{}
				propService.
					On("Create", testifyMock.Anything).
					Return(
						factory.GetProperty(), nil,
					)

				return propService
			},
			asserts: func(got interface{}) {
				assert.Nil(t, got)
				assert.Equal(t, http.StatusCreated, c.Writer.Status())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srv := test.mocks()
			ctrl, err := New(srv)
			assert.Nil(t, err)

			test.configReq(c)

			got := ctrl.Create(c)
			test.asserts(got)
		})
	}
}
