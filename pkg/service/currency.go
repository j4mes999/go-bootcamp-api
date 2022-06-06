package service

import (
	"errors"
	"log"
	"luis/goapi/pkg/client"
	"luis/goapi/pkg/entity"
	"luis/goapi/pkg/repository"
)

const FILE_PATH = "currencies.csv"

type CurrencyService interface {
	CreateCurrencies() error
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
func (cs *currencyService) CreateCurrencies() error {
	log.Printf("SERVICE: CreateCurrencies")
	currencies, err := client.CallExternalAPI()
	if err != nil {
		log.Printf("SERVICE: error in service call")
		return errors.New(err.Error())
	}
	repo := repository.NewCurrencyRepo(FILE_PATH)
	er := repo.WriteCurrencies(currencies)
	if er != nil {
		log.Printf("SERVICE: error in repository")
		return errors.New(err.Error())
	}

	return nil
}

//TODO delete after refactoring
func createObjects(currencies []entity.Currency) error {
	return nil
}
