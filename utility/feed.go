package utility

import (
	"context"

	"github.com/mmcdole/gofeed"
)

func ParseFeed(ctx context.Context, rssStr string) (feed *gofeed.Feed) {

	var (
		err error
		fp  *gofeed.Parser
	)
	fp = gofeed.NewParser()
	feed, err = fp.ParseString(rssStr)
	if err != nil {
		return nil
	}
	return
}
