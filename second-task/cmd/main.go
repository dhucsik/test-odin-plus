package main

import (
	"fmt"
	"second-task/client"
	"second-task/model"
)

const URL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"

func main() {
	var currencies []model.Currency

	cli := client.New()
	cli.GetRequest(&currencies, URL)

	name := "Tether"

	result := cli.FilterByName(currencies, name)

	for _, currency := range result {
		fmt.Println(currency)
	}
}
