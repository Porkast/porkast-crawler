// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2023-08-06 17:07:07
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserSubKeyword is the golang structure for table user_sub_keyword.
type UserSubKeyword struct {
	Id          string      `json:"id"          ` //
	UserId      string      `json:"userId"      ` //
	Keyword     string      `json:"keyword"     ` //
	OrderByDate int         `json:"orderByDate" ` //
	CreateTime  *gtime.Time `json:"createTime"  ` //
	Lang        string      `json:"lang"        ` // feed language
	Status      int         `json:"status"      ` //
}
