package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"luis/goapi/pkg/entity"
	"net/http"
)

const API_ENDPOINT = "https://api.coinbase.com/v2/currencies"

type Response struct {
	Data []entity.Currency `json:"data"`
}

func CallExternalAPI() ([]entity.Currency, error) {
	response, err := http.Get(API_ENDPOINT)

	if err != nil {
		return nil, errors.New("Error invoking the external service: " + err.Error())
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Error in the response of the external service: " + response.Status)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error reading the response in the external service: " + err.Error())
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	return responseObject.Data, nil

}
