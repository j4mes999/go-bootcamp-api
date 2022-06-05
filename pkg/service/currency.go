package service

import (
	"log"
	"luis/goapi/pkg/entity"
)

type CurrencyService interface {
	CreateCurrencies(currencies []entity.Currency) error
}

type CurrencyRepo interface {
	WriteCurrencies(currencies []entity.Currency) error
}

type currencyService struct {
	repo CurrencyRepo
}

func NewCurrencyService(repo CurrencyRepo) CurrencyService {
	return &currencyService{repo}
}

// ******************** Implementation ********************\
func (cs *currencyService) CreateCurrencies(currencies []entity.Currency) error {
	log.Printf("SERVICE: CreateCurrencies")
	cs.repo.WriteCurrencies(currencies)
	return nil

}

func createObjects(currencies []entity.Currency) error {
	return nil
}
