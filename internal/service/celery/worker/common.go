package worker

import (
	"guoshao-fm-crawler/internal/model/entity"
	"strconv"

	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mmcdole/gofeed"
)

func isStringXml(respStr string) bool {
	var (
		err error
	)
	if respStr != "" {
		_, err = gjson.LoadXml(respStr)
		if err != nil {
			return false
		}
		return true
	}

	return false
}

func feedChannelToModel(uid string, feed gofeed.Feed) (model entity.FeedChannel) {

	model = entity.FeedChannel{
		Id:          uid,
		Title:       feed.Title,
		ChannelDesc: feed.Description,
		ImageUrl:    feed.Image.URL,
		Link:        feed.Link,
		FeedLink:    feed.FeedLink,
		Copyright:   feed.Copyright,
		Language:    feed.Language,
		Author:      feed.Author.Name,
		OwnerName:   feed.ITunesExt.Owner.Name,
		OwnerEmail:  feed.ITunesExt.Owner.Email,
		FeedType:    feed.FeedType,
		Categories:  gstr.Join(feed.Categories, ","),
	}

	return
}

func feedItemToModel(channelId string, item gofeed.Item) (model entity.FeedItem) {

	var itemID = strconv.FormatUint(ghash.RS64([]byte(item.Link+item.Title)), 32)

	if item.ITunesExt == nil {
		return
	}

	model = entity.FeedItem{
		Id:          itemID,
		ChannelId:   channelId,
		Title:       item.Title,
		Link:        item.Link,
		PubDate:     gtime.NewFromTime(*item.PublishedParsed),
		InputDate:   gtime.Now(),
		Duration:    item.ITunesExt.Duration,
		Episode:     item.ITunesExt.Episode,
		EpisodeType: item.ITunesExt.EpisodeType,
		Season:      item.ITunesExt.Season,
		Description: item.Description,
	}

	if item.Image != nil {
		model.ImageUrl = item.Image.URL
	}

	if item.Authors != nil || len(item.Authors) > 0 {
		var (
			authors []string
		)
		for _, author := range item.Authors {
			authors = append(authors, author.Name)
		}
		if len(authors) == 0 {
			model.Author = authors[0]
		} else {
			model.Author = gstr.Join(authors, ",")
		}
	}

	if len(item.Enclosures) > 0 && item.Enclosures[0] != nil {
		model.EnclosureUrl = item.Enclosures[0].URL
		model.EnclosureType = item.Enclosures[0].Type
		model.EnclosureLength = item.Enclosures[0].Length
	}

	return
}