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

// internalFeedChannelDao is internal type for wrapping internal DAO implements.
type internalFeedChannelDao = *internal.FeedChannelDao

// feedChannelDao is the data access object for table feed_channel.
// You can define custom methods on it to extend its functionality as you wish.
type feedChannelDao struct {
	internalFeedChannelDao
}

var (
	// FeedChannel is globally public accessible object for table feed_channel operations.
	FeedChannel = feedChannelDao{
		internal.NewFeedChannelDao(),
	}
)

// Fill with you ideas below.
func InsertFeedChannelIfNotExist(ctx context.Context, model entity.FeedChannel) error {
	var (
		err    error
		result gdb.Record
	)
	result, _ = FeedChannel.Ctx(ctx).Where("id=?", model.Id).One()
	if result.IsEmpty() {
		g.Log().Line().Debugf(ctx, "Insert feed channel %s to DB", model.Title)
		_, err = FeedChannel.Ctx(ctx).Save(model)
	} else {
		return errors.New("The Feed Channel is exist")
	}
	return err
}

func InsertOrUpdateFeedChannel(ctx context.Context, model entity.FeedChannel) error {
	var (
		err    error
	)
	_, err = FeedChannel.Ctx(ctx).Save(model)
	return err
}

func GetAllFeedChannelCount(ctx context.Context) (count int, err error) {
	count, err = FeedChannel.Ctx(ctx).Count()
	if err != nil {
		g.Log().Line().Error(ctx, "Get channel total count failed : ", err)
	}

	return
}

func GetFeedChannelList(ctx context.Context, offset, limit int) (channelModelList []entity.FeedChannel) {
	var (
		err error
	)
	err = FeedChannel.Ctx(ctx).Offset(offset).Limit(limit).Scan(&channelModelList)
	if err != nil {
		g.Log().Line().Error(ctx, "Get feed channel list failed : ", err)
	}
	return
}