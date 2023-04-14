package elasticsearch

import (
	"context"
	"guoshao-fm-crawler/internal/model/entity"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/olivere/elastic/v7"
)

const feedItemMapping = `
{
    "settings": {
        "number_of_shards": 1,
        "number_of_replicas": 0
    },
    "mappings": {
        "properties": {
            "title": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            },
            "author": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            },
            "description": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            },
            "created": {
                "enabled": false
            },
            "id": {
                "enabled": false
            },
            "channelId": {
                "enabled": false
            },
            "link": {
                "enabled": false
            },
            "pubDate": {
                "enabled": false
            },
            "imageUrl": {
                "enabled": false
            },
            "enclosureUrl": {
                "enabled": false
            },
            "enclosureType": {
                "enabled": false
            },
            "enclosureLength": {
                "enabled": false
            },
            "duration": {
                "enabled": false
            },
            "episode": {
                "enabled": false
            },
            "explicit": {
                "enabled": false
            },
            "season": {
                "enabled": false
            },
            "episodeType": {
                "enabled": false
            }
        }
    }
}
`

func (c *GSElastic) CreateFeedItemIndexIfNotExit(ctx context.Context) {
	exists, err := c.Client.IndexExists("feed_item").Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := c.Client.CreateIndex("feed_item").BodyString(feedItemMapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
		}
	}

}

func (c *GSElastic) InsertFeedItemList(ctx context.Context, feedChannel entity.FeedChannel, feedItemList []entity.FeedItem) {
	if len(feedItemList) == 0 {
		return
	}
	bulkRequest := c.Client.Bulk()
	for _, feedItem := range feedItemList {
		esFeedItem := entity.FeedItemESData{}
		gconv.Struct(feedItem, &esFeedItem)
		esFeedItem.ChannelImageUrl = feedChannel.ImageUrl
		esFeedItem.ChannelTitle = feedChannel.Title
		esFeedItem.FeedLink = feedChannel.FeedLink
		indexReq := elastic.NewBulkIndexRequest().Index("feed_item").Id(feedItem.Id).Doc(esFeedItem)
		bulkRequest.Add(indexReq)
	}
	resp, err := bulkRequest.Do(ctx)
	if err != nil || resp.Errors {
		respStr := gjson.New(resp)
		g.Log().Line().Errorf(ctx, "bulk index request failed\nError message : %s \nResponse : %s", err, respStr)
	}
}

func (c *GSElastic) InsertFeedItem(ctx context.Context, feedItem entity.FeedItemESData) (err error) {
	g.Log().Line().Debugf(ctx, "Insert feed item %s to elasticsearch", feedItem.Title)
	esFeedItem := entity.FeedItemESData{}
	gconv.Struct(feedItem, &esFeedItem)
	rootDocs := soup.HTMLParse(feedItem.Description)
	esFeedItem.Description = rootDocs.FullText()
	_, err = elastic.NewIndexService(c.Client).Index("feed_item").Id(feedItem.Id).BodyJson(esFeedItem).Do(ctx)
	if err != nil {
		g.Log().Line().Errorf(ctx, "Insert feed item %s to elasticsearch failed %s", feedItem.Title, err)
		return
	}

	return
}
