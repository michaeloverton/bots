package main

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/michaeloverton/bots/internal/env"
	"github.com/michaeloverton/bots/internal/markov"
	"github.com/michaeloverton/bots/internal/texts"
)

func main() {
	// Load environment vars.
	env, err := env.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Create Twitter client.
	config := oauth1.NewConfig(env.ApostleBot.ApiKey, env.ApostleBot.ApiSecret)
	token := oauth1.NewToken(env.ApostleBot.AccessToken, env.ApostleBot.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	twitterClient := twitter.NewClient(httpClient)

	// Test tweet text.
	text, err := markov.OrderTwo(texts.Revelation, 20)
	if err != nil {
		panic(err)
	}

	// Send a Tweet
	_, _, err = twitterClient.Statuses.Update(text, nil)
	if err != nil {
		panic(err)
	}
}
