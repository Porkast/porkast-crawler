package elasticsearch

import (
	"context"
	"guoshao-fm-crawler/internal/model/entity"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/olivere/elastic/v7"
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

func CreateFeedChannelIndexIfNotExit(ctx context.Context) {
	exists, err := esClient.IndexExists("feed_channel").Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := esClient.CreateIndex("feed_channel").BodyString(feedChannelMapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
		}
	}

}

func InsertFeedChannel(ctx context.Context, feedChannel entity.FeedChannel) {
	bulkRequest := esClient.Bulk()
	esFeedChannel := entity.FeedChannelESData{}
	gconv.Struct(feedChannel, &esFeedChannel)
	indexReq := elastic.NewBulkIndexRequest().Index("feed_channel").Id(feedChannel.Id).Doc(esFeedChannel)
	bulkRequest.Add(indexReq)
	resp, err := bulkRequest.Do(ctx)
	if err != nil || resp.Errors {
		respStr := gjson.New(resp)
		g.Log().Line().Errorf(ctx, "feed channel index request failed\nError message : %s \nResponse : %s", err, respStr)
	}
}
