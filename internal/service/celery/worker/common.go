package worker

import (
	"context"
	"guoshao-fm-crawler/internal/model/entity"
	"guoshao-fm-crawler/internal/service/internal/dao"
	"guoshao-fm-crawler/utility"
	"strconv"

	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mmcdole/gofeed"
)

func isStringRSSXml(respStr string) bool {
	var (
		err  error
		fp   *gofeed.Parser
		feed *gofeed.Feed
	)
	if respStr != "" {
		fp = gofeed.NewParser()
		feed, err = fp.ParseString(respStr)
		if err != nil || feed == nil {
			return false
		}
		return true
	}

	return false
}

func storeFeed(ctx context.Context, respStr string) {
	var (
		feed *gofeed.Feed
	)
	feed = utility.ParseFeed(ctx, respStr)
	if feed != nil {
		var (
			feedChannelMode entity.FeedChannel
			feedItemList    []entity.FeedItem
			feedID          string
		)
		feedID = strconv.FormatUint(ghash.RS64([]byte(feed.Description+feed.Title)), 32)
		feedChannelMode = feedChannelToModel(feedID, *feed)
		for _, item := range feed.Items {
			var (
				feedItem entity.FeedItem
			)
			feedItem = feedItemToModel(feedID, *item)
			feedItemList = append(feedItemList, feedItem)
		}
		dao.InsertFeedChannelIfNotExist(ctx, feedChannelMode)
		for _, item := range feedItemList {
			dao.InsertFeedItemIfNotExist(ctx, item)
		}
	}
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
