package repository

import (
	"errors"
	"log"
	"luis/goapi/pkg/entity"
	"os"
)

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
func (cr *currencyRepo) WriteCurrencies(currencies []entity.Currency) error {
	f, err := os.Create(cr.filePath)
	if err != nil {
		return errors.New("Error openning the file for writing: " + err.Error())
	}
	defer f.Close()
	for _, c := range currencies {
		_, err := f.WriteString(c.ID + "," + c.Name + "\n")
		if err != nil {
			log.Printf("REPOSITORY: Error writing line")
		}
	}
	f.Sync()
	return nil
}
