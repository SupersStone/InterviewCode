package indexer

import "dao-exchange/internal/models"

// ProcessLog 事件处理日志
type ProcessLog struct {
	models.Base
	EventBase
	Status  int    `gorm:"column:status;type:tinyint(4) unsigned;comment:1: 成功 ,2:失败" json:"status"`
	LogJSON string `gorm:"column:log_json;type:text;comment:event json" json:"log_json"`
	Msg     string `gorm:"column:msg;type:varchar(300);comment:处理消息" json:"msg"`
}

// TableName table name
func (m *ProcessLog) TableName() string {
	return "process_log"
}
