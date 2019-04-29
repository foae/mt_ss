package main

import (
	"fmt"
	"github.com/turnage/graw/reddit"
	"log"
)

type mtssBot struct {
	bot reddit.Bot
}

// Post implements the Post interface.
func (mt *mtssBot) Post(p *reddit.Post) error {
	if p.Deleted || p.Hidden || p.IsSelf || p.Locked {
		log.Printf("Post (ID: %v) - (%v) discarded.", p.ID, p.Title)
		return nil
	}

	if err := mt.bot.Reply(p.Name, string(generate(blob))); err != nil {
		return fmt.Errorf("could not post a reply to Post (ID: %v) - (%v): %v", p.ID, p.Title, err)
	}

	return nil
}
