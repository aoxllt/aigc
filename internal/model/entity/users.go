// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id   int         `json:"id"   orm:"id"   ` //
	Uid  string      `json:"uid"  orm:"uid"  ` //
	Time *gtime.Time `json:"time" orm:"time" ` //
}
