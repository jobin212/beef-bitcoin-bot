package clients

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// BeefClient for getting Bitcoin price
type BeefClient struct {
}

// BeefPrice is the price of beef at a certain time
type BeefPrice struct {
	Date  string
	Price float64
}

// GetBeefPrice returns the current bitcoin price
func (c *BeefClient) GetBeefPrice() (float64, error) {
	csvFile, _ := os.Open("data/Latest_PBEEFUSDM.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var prices []BeefPrice

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		price, _ := strconv.ParseFloat(line[1], 32)
		prices = append(prices, BeefPrice{
			Date:  line[0],
			Price: price / 100,
		})
	}

	return prices[0].Price, nil
}
