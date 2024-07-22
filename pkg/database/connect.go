package database

import (
	"fmt"
	"net/url"

	"github.com/go-logr/logr"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"product-management-system/pkg/config"
)

// InitDatabase ...
func InitDatabase(log logr.Logger, dbConfig config.DatabaseConfig) *gorm.DB {
	log.Info("start to init database")
	charset := "latin1"
	validateDatabaseConfig(&dbConfig)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=%s",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Dbname, charset, url.QueryEscape("Asia/Shanghai"))
	log.Info("build dsn", "dsn", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)

	}
	// 为 `User` 创建表

	// 将 "ENGINE=InnoDB" 添加到创建 `User` 的 SQL 里去
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	return db
}

func validateDatabaseConfig(cfg *config.DatabaseConfig) {
	if cfg.Username == "" {
		cfg.Username = "root"
	}
	if cfg.Password == "" {
		cfg.Password = "default-password"
	}
	if cfg.Dbname == "" {
		cfg.Dbname = "product-management-system-db"
	}
	if cfg.Host == "" {
		cfg.Host = "127.0.0.1"
	}
	if cfg.Port == 0 {
		cfg.Port = 3306
	}

}

// GenerateUUID ...
func GenerateUUID() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
