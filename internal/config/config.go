package config

import (
	"os"
	"strconv"
)

type Config struct {
	SRV_PORT    string `json:"srv_port"`
	WEB_UI      bool   `json:"web_ui"`
	OpenBrowser bool   `json:"open_browser"`
	SECRET      string `json:"secret"`

	DBConfig `json:"db_config"`
}

type DBConfig struct {
	DB_DRIVE string `json:"db_drive"`
	DB_HOST  string `json:"db_host"`
	DB_PORT  string `json:"db_port"`
	DB_USER  string `json:"db_user"`
	DB_PASS  string `json:"db_pass"`
	DB_NAME  string `json:"db_name"`
	DB_DSN   string `json:"-"`
}

func NewConfig(config *Config) *Config {

	local_conf := defaultConfig()

	if config == nil || config.SRV_PORT != "" {
		local_conf = config
	}

	SRV_PORT := os.Getenv("SRV_PORT")
	if SRV_PORT != "" {
		local_conf.SRV_PORT = SRV_PORT
	}

	SRV_WEB_UI := os.Getenv("SRV_WEB_UI")
	if SRV_WEB_UI != "" {
		local_conf.WEB_UI, _ = strconv.ParseBool(SRV_WEB_UI)
	}

	SRV_DB_DRIVE := os.Getenv("SRV_DB_DRIVE")
	if SRV_DB_DRIVE != "" {
		local_conf.DBConfig.DB_DRIVE = SRV_DB_DRIVE
	}

	SRV_DB_HOST := os.Getenv("SRV_DB_HOST")
	if SRV_DB_HOST != "" {
		local_conf.DBConfig.DB_HOST = SRV_DB_HOST
	}

	SRV_DB_PORT := os.Getenv("SRV_DB_PORT")
	if SRV_DB_PORT != "" {
		local_conf.DBConfig.DB_PORT = SRV_DB_PORT
	}

	SRV_DB_USER := os.Getenv("SRV_DB_USER")
	if SRV_DB_USER != "" {
		local_conf.DBConfig.DB_USER = SRV_DB_USER
	}

	SRV_DB_PASS := os.Getenv("SRV_DB_PASS")
	if SRV_DB_PASS != "" {
		local_conf.DBConfig.DB_PASS = SRV_DB_PASS
	}

	SRV_DB_NAME := os.Getenv("SRV_DB_NAME")
	if SRV_DB_NAME != "" {
		local_conf.DBConfig.DB_NAME = SRV_DB_NAME
	}

	SRV_SECRET := os.Getenv("SECRET")
	if SRV_SECRET != "" {
		local_conf.SECRET = SRV_SECRET
	}
	return local_conf
}

func defaultConfig() *Config {
	default_config := Config{
		SRV_PORT:    "8080",
		WEB_UI:      true,
		OpenBrowser: true,
		DBConfig: DBConfig{
			DB_DRIVE: "mysql",
			DB_HOST:  "localhost",
			DB_PORT:  "3306",
			DB_USER:  "",
			DB_PASS:  "",
			DB_NAME:  "",
		},
	}
	return &default_config
}
