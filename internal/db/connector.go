package db

import (
	"fmt"
	"time"

	"github.com/hughiep/tomato-payment-service/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbURL(c *config.Config) string {
	// Get the database URL
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Database.MySqlDbUser,
		c.Database.MySqlDbPassword,
		c.Database.MysqlDbHost,
		c.Database.MysqlDbPort,
		c.Database.MysqlDbName,
	)
}

func Connect(c *config.Config) *gorm.DB {
	// Connect to the database
	dsn := dbURL(c)
	connection := mysql.Open(dsn)
	mysql, err := gorm.Open(connection, &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	rawDB, err := mysql.DB()
	if err != nil {
		panic("Failed to get raw database")
	}

	// Set the maximum number of open connections
	rawDB.SetMaxOpenConns(10)

	// Set the maximum number of idle connections
	rawDB.SetMaxIdleConns(10)

	// Set the maximum lifetime of a connection
	rawDB.SetConnMaxLifetime(time.Hour)

	return mysql
}
