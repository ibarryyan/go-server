package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RDB *redis.Client

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
	add := viper.GetString("redis.url")
	pwd := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")
	RDB = redis.NewClient(&redis.Options{
		Addr:     add,
		Password: pwd,
		DB:       db,
	})
}
