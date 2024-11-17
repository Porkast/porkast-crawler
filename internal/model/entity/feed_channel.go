// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FeedChannel is the golang structure for table feed_channel.
type FeedChannel struct {
	Id          string `json:"id"          orm:"id"           ` //
	Title       string `json:"title"       orm:"title"        ` //
	ChannelDesc string `json:"channelDesc" orm:"channel_desc" ` //
	ImageUrl    string `json:"imageUrl"    orm:"image_url"    ` //
	Link        string `json:"link"        orm:"link"         ` //
	FeedLink    string `json:"feedLink"    orm:"feed_link"    ` //
	Copyright   string `json:"copyright"   orm:"copyright"    ` //
	Language    string `json:"language"    orm:"language"     ` //
	Author      string `json:"author"      orm:"author"       ` //
	OwnerName   string `json:"ownerName"   orm:"owner_name"   ` //
	OwnerEmail  string `json:"ownerEmail"  orm:"owner_email"  ` //
	FeedType    string `json:"feedType"    orm:"feed_type"    ` //
	Categories  string `json:"categories"  orm:"categories"   ` //
	Source      string `json:"source"      orm:"source"       ` //
	FeedId      string `json:"feedId"      orm:"feed_id"      ` //
}
