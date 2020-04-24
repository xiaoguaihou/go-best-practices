package core

import (
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron/v3"
)

const (
	STATUS_IDEL = iota
	STATUS_STARTED
)

type CoreService struct {
	salesDB *gorm.DB
	status  uint8
	cron    *cron.Cron
}

// service should be created by New function
func New(sldb *gorm.DB) *CoreService {
	return &CoreService{
		salesDB: sldb,
		status:  STATUS_IDEL,
	}
}

// service is started by Start() function
func (s *CoreService) Start() {
	if s.status == STATUS_IDEL {
		c := cron.New()
		_, err := c.AddFunc("0 * * * *", func() {
			s.sendHourReport()
		})
		if err == nil {
			c.Start()
			s.cron = c
		}
	}
}
