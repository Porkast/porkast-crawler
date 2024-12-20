// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FeedChannelDao is the data access object for table feed_channel.
type FeedChannelDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns FeedChannelColumns // columns contains all the column names of Table for convenient usage.
}

// FeedChannelColumns defines and stores column names for table feed_channel.
type FeedChannelColumns struct {
	Id          string //
	Title       string //
	ChannelDesc string //
	ImageUrl    string //
	Link        string //
	FeedLink    string //
	Copyright   string //
	Language    string //
	Author      string //
	OwnerName   string //
	OwnerEmail  string //
	FeedType    string //
	Categories  string //
	Source      string //
	FeedId      string //
}

// feedChannelColumns holds the columns for table feed_channel.
var feedChannelColumns = FeedChannelColumns{
	Id:          "id",
	Title:       "title",
	ChannelDesc: "channel_desc",
	ImageUrl:    "image_url",
	Link:        "link",
	FeedLink:    "feed_link",
	Copyright:   "copyright",
	Language:    "language",
	Author:      "author",
	OwnerName:   "owner_name",
	OwnerEmail:  "owner_email",
	FeedType:    "feed_type",
	Categories:  "categories",
	Source:      "source",
	FeedId:      "feed_id",
}

// NewFeedChannelDao creates and returns a new DAO object for table data access.
func NewFeedChannelDao() *FeedChannelDao {
	return &FeedChannelDao{
		group:   "default",
		table:   "feed_channel",
		columns: feedChannelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FeedChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FeedChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FeedChannelDao) Columns() FeedChannelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FeedChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FeedChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FeedChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
