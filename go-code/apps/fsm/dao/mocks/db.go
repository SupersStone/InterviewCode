package mocks

import "dao-exchange/pkg/orm"

// BuildDBCfg mock db
func BuildDBCfg() orm.DBCfg {
	return orm.DBCfg{
		Dsn:         "root:meta@tcp(localhost:3306)/nft_local?timeout=2s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4",
		MaxOpenConn: 3,
		MaxIdleConn: 10,
		LogLevel:    4,
	}
}
