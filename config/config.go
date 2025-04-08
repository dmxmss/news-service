package config

import (
	"github.com/spf13/viper"

	"sync"
	"strings"
)

type (
	Config struct {
		App *App
		Database *Database
	}

	App struct {
		Address string
		Port string
		JwtSecret string
	}

	Database struct {
		Host string
		Name string	
		User string
		Port string
		Password string
	}
)

var (
	once sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func () {
		viper.AddConfigPath("./")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")	

		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		viper.SetDefault("app.address", "localhost")
		viper.SetDefault("app.port", "8000")

		viper.SetDefault("database.name", "postrges")
		viper.SetDefault("database.user", "postgres")
		viper.SetDefault("database.port", "5432")
		viper.SetDefault("database.host", "db")

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
