package clients

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const bitcoinPriceURL = "https://blockchain.info/tobtc?currency=USD&value=10.99"

// BitcoinClient for getting Bitcoin price
type BitcoinClient struct {
}

// GetBitcoinPrice returns the current bitcoin price
func (c *BitcoinClient) GetBitcoinPrice() (float64, error) {
	resp, err := http.Get(bitcoinPriceURL)
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return strconv.ParseFloat(string(body), 32)
}
