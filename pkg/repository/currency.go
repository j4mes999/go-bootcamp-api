package repository

import "luis/goapi/pkg/entity"

type CurrencyRepo interface {
	WriteCurrencies(currencies []entity.Currency) error
}

type currencyRepo struct {
	filePath string
}

func NewCurrencyRepo(file string) CurrencyRepo {
	return &currencyRepo{file}
}

// ******************** Implementation ********************
func (*currencyRepo) WriteCurrencies(currencies []entity.Currency) error {
	return nil
}
