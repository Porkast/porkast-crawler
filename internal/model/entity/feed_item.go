// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FeedItem is the golang structure for table feed_item.
type FeedItem struct {
	Id              string      `json:"id"              orm:"id"               ` //
	ChannelId       string      `json:"channelId"       orm:"channel_id"       ` //
	Guid            string      `json:"guid"            orm:"guid"             ` //
	Title           string      `json:"title"           orm:"title"            ` //
	Link            string      `json:"link"            orm:"link"             ` //
	PubDate         *gtime.Time `json:"pubDate"         orm:"pub_date"         ` //
	Author          string      `json:"author"          orm:"author"           ` //
	InputDate       *gtime.Time `json:"inputDate"       orm:"input_date"       ` //
	ImageUrl        string      `json:"imageUrl"        orm:"image_url"        ` //
	EnclosureUrl    string      `json:"enclosureUrl"    orm:"enclosure_url"    ` //
	EnclosureType   string      `json:"enclosureType"   orm:"enclosure_type"   ` //
	EnclosureLength string      `json:"enclosureLength" orm:"enclosure_length" ` //
	Duration        string      `json:"duration"        orm:"duration"         ` //
	Episode         string      `json:"episode"         orm:"episode"          ` //
	Explicit        string      `json:"explicit"        orm:"explicit"         ` //
	Season          string      `json:"season"          orm:"season"           ` //
	Episodetype     string      `json:"episodetype"     orm:"episodetype"      ` //
	Description     string      `json:"description"     orm:"description"      ` //
	ChannelTitle    string      `json:"channelTitle"    orm:"channel_title"    ` //
	FeedId          string      `json:"feedId"          orm:"feed_id"          ` //
	FeedLink        string      `json:"feedLink"        orm:"feed_link"        ` //
	Source          string      `json:"source"          orm:"source"           ` //
}
