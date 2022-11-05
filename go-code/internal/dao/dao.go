package dao

import (
	"log"

	"dao-exchange/pkg/orm"

	"gorm.io/gorm"
)

// BaseDao db about
type BaseDao struct {
	db *gorm.DB
}

// NewDaoWithDB new dao for tx
func NewDaoWithDB(db *gorm.DB) *BaseDao {
	return &BaseDao{
		db: db,
	}
}

// NewDao new dao
func NewDao(cfg orm.DBCfg) *BaseDao {
	db, err := orm.NewGorm(cfg)
	if err != nil {
		log.Fatalf("connect db err %s", err)
	}

	return &BaseDao{db: db}
}

// Close close db
func (d *BaseDao) Close() {
	sqlDB, err := d.db.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
}

// Db  db
func (d *BaseDao) Db() *gorm.DB {
	return d.db
}
