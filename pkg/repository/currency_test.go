package repository

import (
	"luis/goapi/pkg/entity"
	"luis/goapi/test/testData"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrencyRepo_WriteCurrencies(t *testing.T) {

	var testCases = []struct {
		name            string
		paramCurrencies []entity.Currency
		error
	}{
		{
			"Should be succesful",
			testData.Currencies,
			nil,
		}, //Not very sure how to add more test cases
		/* 	{
			"Should not be succesful",
			testData.Currencies,
			errors.New("Error invoking the external API"),
		}, */
	}

	repo := NewCurrencyRepo("currencies_test.csv")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.WriteCurrencies(tc.paramCurrencies)
			assert.Equal(t, err, tc.error)
		})
	}
}
