package main

import (
	"fmt"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type mtssBot struct {
	bot reddit.Bot
}

func main() {
	subreddit := mustGetEnv("SUBREDDIT")
	credsFile := mustGetEnv("REDDIT_AGENT_CREDENTIALS_FILE_NAME")
	//subreddit = "romania_ss"

	buf, err := ioutil.ReadFile(credsFile)
	if err != nil {
		log.Fatalf("could not read file (%v) to read: %v", credsFile, err)
	}
	log.Println("-------------")
	log.Printf("\n%s\n", buf)
	log.Println("-------------")

	if bot, err := reddit.NewBotFromAgentFile(credsFile, time.Minute*2); err != nil {
		log.Printf("could not create bot handle: %v", err)
		log.Printf("%#v", bot)
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
		if err := wait(); err != nil {
			log.Fatalf("could not wait: %v", err)
		}
	}

}

// Post implements the Post interface.
func (mt *mtssBot) Post(p *reddit.Post) error {
	if p.Deleted || p.Hidden || p.IsSelf || p.Locked {
		log.Printf("Post (ID: %v) - (%v) discarded.", p.ID, p.Title)
		return nil
	}

	if err := mt.bot.Reply(p.Name, "das tecst"); err != nil {
		return fmt.Errorf("could not post a reply to Post (ID: %v) - (%v): %v", p.ID, p.Title, err)
	}

	return nil
}

func mustGetEnv(value string) string {
	v := os.Getenv(value)
	if v == "" {
		log.Fatalf("could not retrieve needed value (%v) from the environment", value)
	}

	return v
}
