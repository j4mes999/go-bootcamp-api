package service

import (
	_ "errors"
	"luis/goapi/pkg/entity"
	"luis/goapi/test/testData"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrencyService_CreateCurrencies(t *testing.T) {
	var testCases = []struct {
		name            string
		paramCurrencies []entity.Currency
		error
	}{
		{
			"Should be succesful",
			testData.Currencies,
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := createObjects(tc.paramCurrencies)
			assert.Equal(t, err, tc.error)
		})
	}
}
