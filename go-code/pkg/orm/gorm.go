package orm

import (
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DBCfg database config
type DBCfg struct {
	Dsn         string `json:"dsn"`
	MaxOpenConn int    `json:"max_open_conn"`
	MaxIdleConn int    `json:"max_idle_conn"`
	LogLevel    int    `json:"log_level"`
}

// NewGorm new dao
func NewGorm(cfg DBCfg) (db *gorm.DB, err error) {
	slowLogger := logger.New(
		//设置Logger
		NewDBLog(),
		logger.Config{
			SlowThreshold: time.Millisecond * 100,
			LogLevel:      logger.LogLevel(cfg.LogLevel),
		},
	)

	db, err = gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger:                                   slowLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		ConnPool:                                 nil,
		Dialector:                                nil,
	})
	if err != nil {
		err = errors.Wrap(err, "connect to db error")
		return
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = sqlDB.Ping()

	return
}
