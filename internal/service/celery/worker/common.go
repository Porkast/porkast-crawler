package worker

import (
	"context"
	"fmt"
	"guoshao-fm-crawler/internal/model/entity"
	"guoshao-fm-crawler/internal/service/elasticsearch"
	"guoshao-fm-crawler/internal/service/internal/dao"
	"guoshao-fm-crawler/utility"
	"strconv"

	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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

func storeFeed(ctx context.Context, respStr, feedLink string) {
	var (
		feed *gofeed.Feed
	)
	feed = utility.ParseFeed(ctx, respStr)
	defer func(ctx context.Context) {
		if rec := recover(); rec != nil {
			g.Log().Line().Error(ctx, fmt.Sprintf("Store feed failed:\n%s\nThe feed string :\n%s\n", rec, respStr))
		}
	}(ctx)
	if feed != nil {
		var (
			err             error
			feedChannelMode entity.FeedChannel
			feedItemList    []entity.FeedItem
			feedID          string
		)
		feedID = strconv.FormatUint(ghash.RS64([]byte(feed.FeedLink+feed.Title)), 32)
		feedChannelMode = feedChannelToModel(feedID, *feed)
		if feedChannelMode.FeedLink == "" {
			feedChannelMode.FeedLink = feedLink
		}
		for _, item := range feed.Items {
			var (
				feedItem entity.FeedItem
			)
			feedItem = feedItemToModel(feedID, *item)
			if feedItem.Author == "" {
				feedItem.Author = feedChannelMode.Author
			}
			feedItemList = append(feedItemList, feedItem)
		}
		err = dao.InsertOrUpdateFeedChannel(ctx, feedChannelMode)
		if err == nil {
			elasticsearch.Client().InsertFeedChannel(ctx, feedChannelMode)
		}
		for _, item := range feedItemList {
			err = dao.InsertFeedItemIfNotExist(ctx, item)
			if err == nil {
				var esFeedItem entity.FeedItemESData
				gconv.Struct(item, &esFeedItem)
				esFeedItem.ChannelImageUrl = feedChannelMode.ImageUrl
				esFeedItem.ChannelTitle = feedChannelMode.Title
				esFeedItem.FeedLink = feedChannelMode.FeedLink
				esFeedItem.Language = feedChannelMode.Language
				elasticsearch.Client().InsertFeedItem(ctx, esFeedItem)
			}
		}
	}
}

func feedChannelToModel(uid string, feed gofeed.Feed) (model entity.FeedChannel) {

	var (
		authorList []string
	)
	model = entity.FeedChannel{
		Id:          uid,
		Title:       feed.Title,
		ChannelDesc: feed.Description,
		Link:        feed.Link,
		FeedLink:    feed.FeedLink,
		Copyright:   feed.Copyright,
		Language:    feed.Language,
		FeedType:    feed.FeedType,
		Categories:  gstr.Join(feed.Categories, ","),
	}

	if len(feed.Authors) > 1 {
		for _, authorItem := range feed.Authors {
			var (
				author string
			)
			author = authorItem.Name
			authorList = append(authorList, author)
		}
		model.Author = gstr.Join(authorList, ",")
	} else if len(feed.Authors) == 1 {
		model.Author = feed.Authors[0].Name
	}

	if feed.ITunesExt != nil && feed.ITunesExt.Owner != nil {
		model.OwnerName = feed.ITunesExt.Owner.Name
		model.OwnerEmail = feed.ITunesExt.Owner.Email
	}

	if feed.Image != nil {
		model.ImageUrl = feed.Image.URL
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
			author.Name = formatFeedAuthor(author.Name)
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

func formatFeedAuthor(author string) (formatAuthor string) {

	if author != "" && gstr.HasSuffix(author, "|") {
		formatAuthor = author[:len(author)-1]
	} else {
		formatAuthor = author
	}

	return
}
