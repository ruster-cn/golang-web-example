package orm

import (
	"fmt"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConnectConfiguration struct {
	User         string          `yaml:"user"`
	Password     string          `yaml:"password"`
	Host         string          `yaml:"host"`
	Port         string          `yaml:"port"`
	DBName       string          `yaml:"db_name"`
	TimeOut      string          `yaml:"time_out"`
	MaxIdleConns int             `yaml:"max_idle_conns"`
	MaxOpenConns int             `yaml:"max_open_conns"`
	MaxLifetime  time.Duration   `yaml:"max_lifetime"`
	LogMode      logger.LogLevel `yaml:"log_mode"`
}

//NewMysqlEngine new gorm mysql engine.engine will be used by Dao
func NewMysqlEngine(config *MysqlConnectConfiguration) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=%s&charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DBName, config.TimeOut)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(config.LogMode),
	})

	if err != nil {
		return nil, fmt.Errorf("open mysql connect fail,%v", err)
	}

	pool, err := db.DB()
	if err != nil {
		return nil, err
	}
	pool.SetConnMaxLifetime(config.MaxLifetime)
	pool.SetMaxIdleConns(config.MaxIdleConns)
	pool.SetMaxOpenConns(config.MaxOpenConns)

	return db, nil
}
