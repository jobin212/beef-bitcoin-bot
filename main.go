package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jobin212/beef-bot/clients"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Go-Twitter Bot v0.01")

	twitterCreds := clients.TwitterCredentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	twitterClient := clients.TwitterClient{}
	bitcoinClient := clients.BitcoinClient{}

	err = twitterClient.InitTwitterClient(&twitterCreds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	price, err := bitcoinClient.GetBitcoinPrice()
	if err != nil {
		log.Println(err)
	}

	message := fmt.Sprintf("1 lb of costco ribeye is approximately: %.10f₿", price)
	log.Printf(message)
	err = twitterClient.UpdateStatus(message)
	if err != nil {
		log.Println(err)
	}
}
