package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginAccountQueryOCO()
}

func MarginAccountQueryOCO() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, binance_connector.SIGNATURE_HMAC_SHA256, baseURL)

	// MarginAccountQueryOCOService - /sapi/v1/margin/orderList
	marginAccountQueryOCO, err := client.NewMarginAccountQueryOCOService().OrigClientOrderId("").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountQueryOCO))
}
