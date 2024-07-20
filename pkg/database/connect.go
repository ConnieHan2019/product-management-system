package database

import (
	"fmt"
	"net/url"
	"os"

	"github.com/go-logr/logr"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"product-management-system/config"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func TESTConnectDB() *gorm.DB {
	dsn := "host=localhost port=5432 dbname=product-management-system-db user=user password=pass sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// InitDatabase ...
func InitDatabase(log logr.Logger, dbConfig config.DatabaseConfig) *gorm.DB {
	log.Info("start to init database", "config")
	charset := "latin1"
	validateDatabaseConfig(&dbConfig)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=%s",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Dbname, charset, url.QueryEscape("Asia/Shanghai"))
	log.Info("build dsn", "dsn", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}
	// 为 `User` 创建表

	// 将 "ENGINE=InnoDB" 添加到创建 `User` 的 SQL 里去
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	return db
}

func validateDatabaseConfig(cfg *config.DatabaseConfig) {
	if cfg.Username == "" {
		cfg.Username = "user"
	}
	if cfg.Password == "" {
		cfg.Password = "Hello@1234"
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
