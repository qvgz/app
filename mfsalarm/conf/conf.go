package conf

import (
	"github.com/qvgz/golib/file"
	"github.com/qvgz/golib/mail"
)

type Conf struct {
	Smtp      mail.Smtp `json:"smtp"`
	DiskAlarm DiskAlarm `json:"diskalarm"`
}

type DiskAlarm struct {
	Url               string
	Mails             []string `json:"mails"`
	CheckIntervalTime string   `json:"checkIntervalTime"`
	AlarmIntervalTime int      `json:"alarmIntervalTime"`
}

func (c *Conf) Init() error {
	return file.JsonInitValue("conf.json", c)
}
