package validator

import (
	"testing"

	infraErrors "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/errors"
	"github.com/stretchr/testify/assert"
)

type (
	DataWithoutValidation struct {
		ID     int    `json:"id"`
		Text   string `json:"text"`
		Total  *uint  `json:"total,omitempty"`
		Hidden string `json:"-"`
	}

	Data struct {
		ID           int     `json:"id" validate:"required,min=5,max=10"`
		Text         string  `json:"text" validate:"required,max=5"`
		Date         string  `json:"date" validate:"required,datetime=2006-01-02"`
		OptionalText *string `json:"optional_text"`
	}
)

func TestValidate(t *testing.T) {
	type testCase struct {
		name    string
		input   interface{}
		asserts func(*testing.T, error)
	}

	tests := []testCase{
		{
			name: "error when have invalid date format",
			input: Data{
				ID:   5,
				Text: ":)",
				Date: "2021/10/30",
			},
			asserts: func(t *testing.T, err error) {
				assert.Equal(t, err, getWantError([]string{"Date does not match the 2006-01-02 format"}))
			},
		},
		{
			name:  "error when dont have required fields ",
			input: Data{},
			asserts: func(t *testing.T, err error) {
				assert.Equal(t, err, getWantError([]string{
					"ID is required",
					"Text is required",
					"Date is required",
				}))
			},
		},
		{
			name:  "max ID value violated",
			input: Data{ID: 55, Text: "hello", Date: "2021-10-30"},
			asserts: func(t *testing.T, err error) {
				assert.Equal(t, err, getWantError([]string{
					"ID must be 10 or less",
				}))
			},
		},
		{
			name:  "max ID value violated",
			input: Data{ID: 6, Text: "im a test", Date: "2021-10-30"},
			asserts: func(t *testing.T, err error) {
				assert.Equal(t, err, getWantError([]string{
					"Text must be a maximum of 5 characters in length",
				}))
			},
		},
		{
			name:  "success when no validation rule",
			input: DataWithoutValidation{},
			asserts: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.asserts(t, Validate(test.input))
		})
	}
}

func getWantError(causes []string) infraErrors.MultiCauseError {
	return infraErrors.NewMultiCauseError("invalid data", causes...)
}
