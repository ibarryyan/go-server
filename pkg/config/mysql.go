package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	viper.SetConfigName("app")
	viper.SetConfigType("properties")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No file ...")
		} else {
			fmt.Println("Find file but have err ...")
		}
	}
	RedisUrl = viper.GetString("redis.url")
	RedisPwd = viper.GetString("redis.password")
	RedisDb = viper.GetInt("redis.db")
	InitRedis()
	PORT = viper.GetString("server.port")
	url := viper.GetString("db.url")
	db := viper.GetString("db.databases")
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	dsn := username + ":" + password + "@tcp(" + url + ")/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
