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
	beefClient := clients.BeefClient{}

	err = twitterClient.InitTwitterClient(&twitterCreds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	beefPrice, err := beefClient.GetBeefPrice()
	if err != nil {
		log.Println(err)
	}

	bitcoinPrice, err := bitcoinClient.GetBitcoinPrice(beefPrice)
	if err != nil {
		log.Println(err)
	}
	satoshiPrice := 100000000 * bitcoinPrice

	cattlePrice, err := bitcoinClient.GetBitcoinPrice(975.0)
	if err != nil {
		log.Println(err)
	}
	numCattle := 1 / cattlePrice

	message := fmt.Sprintf("1lb of beef is approximately %.0f satoshis globally 🌎\nWith one bitcoin you can buy about %.1f Hereford Heiferettes in Wyoming\n🐄🐄🐄🐄🐄🐄🐄🐄🐄",
		satoshiPrice, numCattle)
	log.Printf(message)
	err = twitterClient.UpdateStatus(message)
	if err != nil {
		log.Println(err)
	}
}
