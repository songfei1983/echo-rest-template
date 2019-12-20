package db

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/songfei1983/go-api-server/pkg/logger"
	gormzap "github.com/wantedly/gorm-zap"
)

var global *gorm.DB
var once sync.Once

func New() *gorm.DB {
	once.Do(func() {
		global = connect()
	})
	return global
}

func connect() *gorm.DB{
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.LogMode(true)
	db.SetLogger(gormzap.New(logger.New()))
	return db
}
