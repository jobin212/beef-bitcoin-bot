package clients

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// BitcoinClient for getting Bitcoin price
type BitcoinClient struct {
}

// GetBitcoinPrice returns the current bitcoin price
func (c *BitcoinClient) GetBitcoinPrice(price float64) (float64, error) {

	url := fmt.Sprintf("https://blockchain.info/tobtc?currency=USD&value=%f", price)
	resp, err := http.Get(url)
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := strconv.ParseFloat(string(body), 32)
	if err != nil {
		return -1, err
	}

	return f, nil
}
