package clients

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// TwitterClient for getting Bitcoin price
type TwitterClient struct {
	client *twitter.Client
}

// TwitterCredentials store all of our tokens
type TwitterCredentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// InitTwitterClient initializes twitter client
func (c *TwitterClient) InitTwitterClient(creds *TwitterCredentials) error {

	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)

	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	twitterClient := twitter.NewClient(httpClient)

	// Verify credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	_, _, err := twitterClient.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return err
	}

	c.client = twitterClient

	return nil
}

// UpdateStatus allows user to update twitter bot status
func (c *TwitterClient) UpdateStatus(message string) error {
	_, _, err := c.client.Statuses.Update(message, nil)
	return err
}
