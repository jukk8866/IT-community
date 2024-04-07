package models

import "time"

/*
	type Community struct {
		ID   int64
		Name string
	}
*/
type Community struct {
	CommunityID  int64     //`gorm:"column: community_id "`
	Name         string    //`gorm:"column: name "`
	Introduction string    //`gorm:"column: introduction "`
	CreateTime   time.Time //`gorm:"column: create_time "`
}

/*
	type CommunityDetail struct {
		ID           int64
		Name         string
		Introduction string
		CreateTime   time.Time
	}
*/
/*func (Community) TableName() string {
return "community" //关掉复数就行
*/
