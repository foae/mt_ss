package main

import (
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
	"log"
	"os"
	"time"
)

func main() {
	subreddit := mustGetEnv("SUBREDDIT")
	credsFile := mustGetEnv("REDDIT_AGENT_CREDENTIALS_FILE_NAME")

	if bot, err := reddit.NewBotFromAgentFile(credsFile, time.Minute*2); err != nil {
		log.Printf("could not create bot handle: %v", err)
	} else {
		cfg := graw.Config{
			CommentReplies: false,
			Mentions:       false,
			Messages:       false,
			PostReplies:    false,
			Subreddits:     []string{subreddit},
		}
		handler := &mtssBot{bot: bot}
		_, wait, err := graw.Run(handler, bot, cfg)
		if err != nil {
			log.Fatalf("could not run mtssbot: %v", err)
		}

		log.Println("Waiting patiently...")
		if err := wait(); err != nil {
			log.Printf("could not wait: %v", err)
		}
	}

	// TODO: handle token refresh?

}

func mustGetEnv(value string) string {
	v := os.Getenv(value)
	if v == "" {
		log.Fatalf("could not retrieve needed value (%v) from the environment", value)
	}

	return v
}
