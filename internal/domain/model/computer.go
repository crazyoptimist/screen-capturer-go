package model

import (
	"fmt"
	"screencapturer/internal/constant"
)

type Computer struct {
	Common
	Name      string `gorm:"column:name;unique" json:"name"`
	IsActive  bool   `gorm:"column:is_active;default:false" json:"isActive"`
	IPAddress string `gorm:"column:ip_address;default:''" json:"ipAddress"`
}

func (c *Computer) GetEndpoint() string {
	if c.IPAddress == "" {
		return ""
	}

	return fmt.Sprintf("http://%s:%v", c.IPAddress, constant.SERVER_WEB_PORT)
}
