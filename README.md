API exposing one endpoint:
  GET /api/v1/currencies: It will call the API exposed on:  https://api.coinbase.com/v2/currencies which
  retrieves the data of all currencies in the exchange with an ID and a NAME. Once the data is fetched 
  the API stores this data in a .csv file called "currencies.csv".