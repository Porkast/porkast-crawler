// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FeedItemDao is the data access object for table feed_item.
type FeedItemDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns FeedItemColumns // columns contains all the column names of Table for convenient usage.
}

// FeedItemColumns defines and stores column names for table feed_item.
type FeedItemColumns struct {
	Id              string //
	ChannelId       string //
	Guid            string //
	Title           string //
	Link            string //
	PubDate         string //
	Author          string //
	InputDate       string //
	ImageUrl        string //
	EnclosureUrl    string //
	EnclosureType   string //
	EnclosureLength string //
	Duration        string //
	Episode         string //
	Explicit        string //
	Season          string //
	Episodetype     string //
	Description     string //
	ChannelTitle    string //
	FeedId          string //
	FeedLink        string //
	Source          string //
}

// feedItemColumns holds the columns for table feed_item.
var feedItemColumns = FeedItemColumns{
	Id:              "id",
	ChannelId:       "channel_id",
	Guid:            "guid",
	Title:           "title",
	Link:            "link",
	PubDate:         "pub_date",
	Author:          "author",
	InputDate:       "input_date",
	ImageUrl:        "image_url",
	EnclosureUrl:    "enclosure_url",
	EnclosureType:   "enclosure_type",
	EnclosureLength: "enclosure_length",
	Duration:        "duration",
	Episode:         "episode",
	Explicit:        "explicit",
	Season:          "season",
	Episodetype:     "episodetype",
	Description:     "description",
	ChannelTitle:    "channel_title",
	FeedId:          "feed_id",
	FeedLink:        "feed_link",
	Source:          "source",
}

// NewFeedItemDao creates and returns a new DAO object for table data access.
func NewFeedItemDao() *FeedItemDao {
	return &FeedItemDao{
		group:   "default",
		table:   "feed_item",
		columns: feedItemColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FeedItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FeedItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FeedItemDao) Columns() FeedItemColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FeedItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FeedItemDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FeedItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
