package elasticsearch

import (
	"context"
	"porkast-crawler/internal/model/entity"
)

const feedChannelMapping = `
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
            "ownerName": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            },
            "author": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            },
            "channelDesc": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            },
            "id": {
                "enabled": false
            },
            "link": {
                "enabled": false
            },
            "feedLink": {
                "enabled": false
            },
            "language": {
                "enabled": false
            },
            "copyright": {
                "enabled": false
            },
            "imageUrl": {
                "enabled": false
            },
            "ownerEmail": {
                "enabled": false
            },
            "feedType": {
                "enabled": false
            },
            "categories": {
                "enabled": false
            }
        }
    }
}
`

func (c *GSElastic) CreateFeedChannelIndexIfNotExit(ctx context.Context) {
	exists, err := c.Client.IndexExists("feed_channel").Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := c.Client.CreateIndex("feed_channel").BodyString(feedChannelMapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
		}
	}

}

func (c *GSElastic) InsertFeedChannel(ctx context.Context, feedChannel entity.FeedChannel) {
	// g.Log().Line().Debugf(ctx, "Insert feed channel %s to elasticsearch", feedChannel.Title)
	// bulkRequest := c.Client.Bulk()
	// esFeedChannel := entity.FeedChannelESData{}
	// gconv.Struct(feedChannel, &esFeedChannel)
	// rootDocs := soup.HTMLParse(feedChannel.ChannelDesc)
	// esFeedChannel.TextChannelDesc = rootDocs.FullText()
	// indexReq := elastic.NewBulkIndexRequest().Index("feed_channel").Id(feedChannel.Id).Doc(esFeedChannel)
	// bulkRequest.Add(indexReq)
	// resp, err := bulkRequest.Do(ctx)
	// if err != nil || resp.Errors {
	// 	respStr := gjson.New(resp)
	// 	g.Log().Line().Errorf(ctx, "feed channel index request failed\nError message : %s \nResponse : %s", err, respStr.MustToIniString())
	// }
}
