// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"errors"
	"porkast-crawler/internal/dao/internal"
	"porkast-crawler/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// internalFeedItemDao is internal type for wrapping internal DAO implements.
type internalFeedItemDao = *internal.FeedItemDao

// feedItemDao is the data access object for table feed_item.
// You can define custom methods on it to extend its functionality as you wish.
type feedItemDao struct {
	internalFeedItemDao
}

var (
	// FeedItem is globally public accessible object for table feed_item operations.
	FeedItem = feedItemDao{
		internal.NewFeedItemDao(),
	}
)

// Fill with you ideas below.
func InsertFeedItemIfNotExist(ctx context.Context, model entity.FeedItem) (err error) {

	var (
		result gdb.Record
	)

	result, err = FeedItem.Ctx(ctx).Where("channel_id=?", model.ChannelId).Where("title=?", model.Title).One()
	if err != nil {
		return
	}

	if !result.IsEmpty() {
		return errors.New("The feed item is exist.")
	}

	g.Log().Line().Debugf(ctx, "Insert feed item %s to DB", model.Title)
	_, err = FeedItem.Ctx(ctx).Insert(model)

	return
}
