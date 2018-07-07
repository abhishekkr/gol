package main

import (
	"fmt"
	"log"

	golgoquery "github.com/abhishekkr/gol/golgoquery"
)

var (
	actionmap = map[string]golgoquery.Action{"Tweet": Tweet}

	tweetsQuery = `
{ "uriflows": [
    { "uri": "https://twitter.com/abionic",
      "queryflows": [ {
			"queries": [
			{ "selector": [".fullname.show-popup-with-id"], "attribute": "text", "action": "~" },
			{ "selector": [".js-tweet-text-container"], "attribute": "text", "action": "~" }
			],
			"action": "Tweet"
		}
      ]
	}
  ] }`
)

func main() {
	_, err := golgoquery.FromJson([]byte(tweetsQuery), actionmap)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("~finito")
}

func printTweet(account, tweet string) {
	fmt.Printf("@%s : %s\n", account, tweet)
	fmt.Println("*********************************************************")
}

func Tweet(resultsList []golgoquery.GoqueryResults, idx int) ([]golgoquery.GoqueryResults, error) {
	accounts := resultsList[0].Results
	tweets := resultsList[1].Results

	for idx, account := range accounts {
		printTweet(account, tweets[idx])
	}
	return resultsList, nil
}
