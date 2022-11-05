package repo

import (
	"log"
	"time"

	"dao-exchange/pkg/orm"

	"gorm.io/gorm"
)

var timeout = time.Second * 5

// Dao db about
type Dao struct {
	db *gorm.DB
}

// New new dao
func New(cfg orm.DBCfg) *Dao {
	db, err := orm.NewGorm(cfg)
	if err != nil {
		log.Fatalf("connect db err %s", err)
	}

	return &Dao{
		db: db,
	}
}

// Close close db
func (d *Dao) Close() {
	sqlDB, err := d.db.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
}
