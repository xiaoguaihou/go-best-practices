package core

// database orm struct

import (
	"time"
)

type ConfigRouting struct {
	ID          int       `gorm:"column:id;primary_key"`
	FromCity    string    `gorm:"column:from_city"`
	ToCity      string    `gorm:"column:to_city"`
	Carrier     string    `gorm:"column:carrier"`
	FromDay     int       `gorm:"column:from_day"`
	RetDay      int       `gorm:"column:ret_day"`
	GmtModified time.Time `gorm:"column:gmt_modified"`
	GmtCreate   time.Time `gorm:"column:gmt_create"`
}
