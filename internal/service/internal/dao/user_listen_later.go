// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"porkast-crawler/internal/service/internal/dao/internal"
)

// userListenLaterDao is the data access object for table user_listen_later.
// You can define custom methods on it to extend its functionality as you wish.
type userListenLaterDao struct {
	*internal.UserListenLaterDao
}

var (
	// UserListenLater is globally public accessible object for table user_listen_later operations.
	UserListenLater = userListenLaterDao{
		internal.NewUserListenLaterDao(),
	}
)

// Fill with you ideas below.
