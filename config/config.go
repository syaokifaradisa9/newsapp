package config

import "github.com/spf13/viper"

type App struct{
	Port string `json:"app_port"`
	ENV string `json:"app_env"`

	JWTSecret string `json:"jwt_secret"`
	JWTIssuer string `json:"jwt_issuer"`

}

type PsqlDB struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	DBName string `json:"db_name"`
	DBMaxOpen int `json:"db_max_open"`
	DBMaxIdle int `json:"db_max_idle"`
}

type Config struct {
	App    App
	PsqlDB PsqlDB
}

func NewConfig() *Config {
	return &Config{
		App: App{
			Port: viper.GetString("APP_PORT"),
			ENV: viper.GetString("APP_ENV"),

			JWTSecret: viper.GetString("JWT_SECRET_KEY"),
			JWTIssuer: viper.GetString("JWT_ISSUER"),
		},
		PsqlDB: PsqlDB{
			Host: viper.GetString("DATABASE_HOST"),
			Port: viper.GetString("DATABASE_PORT"),
			User: viper.GetString("DATABASE_USER"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			DBName: viper.GetString("DATABASE_NAME"),
			DBMaxOpen: viper.GetInt("DATABASE_MAX_OPEN_CONNECTION"),
			DBMaxIdle: viper.GetInt("DATABASE_MAX_IDLE_CONNECTION"),
		},
	}
}