package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App       string `validate:"required"`
	AppVer    string `validate:"required"`
	Env       string `validate:"required"`
	Discord   DiscordConfig
	Http      HttpConfig     `validate:"required"`
	Log       LogConfig      `validate:"required"`
	Database  DatabaseConfig `validate:"dive,required"`
	Scheduler SchedulerConfig
	Data      DataConfig `validate:"required"`
}

type DiscordConfig struct {
	WebhookUrl DiscordWebhookUrl
	AvatarUrl  DiscordAvatarUrl
}

type DiscordWebhookUrl struct {
	OnLeave    string
	OnBirthday string
}

type DiscordAvatarUrl struct {
	LeaveAvatar    string
	WorkAvatar     string
	BirthdayAvatar string
}

type DataConfig struct {
	LimitUnprocessedPPS      int `validate:"required"`
	LimitUnprocessedEmployee int `validate:"required"`
}

type SchedulerConfig struct {
	IntervalInSecond string
	MultiTime        string
}

type HttpConfig struct {
	Port         int `validate:"required"`
	WriteTimeout int `validate:"required"`
	ReadTimeout  int `validate:"required"`
}

type LogConfig struct {
	FileLocation    string `validate:"required"`
	FileTDRLocation string `validate:"required"`
	FileMaxSize     int    `validate:"required"`
	FileMaxBackup   int    `validate:"required"`
	FileMaxAge      int    `validate:"required"`
	Stdout          bool   `validate:"required"`
}

type DatabaseConfig struct {
	Host                      string `validate:"required"`
	User                      string `validate:"required"`
	Password                  string `validate:"required"`
	Name                      string `validate:"required"`
	Port                      string `validate:"required"`
	SSLMode                   string `validate:"required"`
	MaxIdleConn               int    `validate:"required"`
	ConnMaxLifetime           int    `validate:"required"`
	MaxOpenConn               int    `validate:"required"`
	TransactionIsolationLevel string
	TransactionReadOnly       bool
}

func (c *Config) LoadConfig(path string) {
	viper.AddConfigPath(".")
	viper.SetConfigName(path)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}

		if strings.HasPrefix(value, "$[") && strings.HasSuffix(value, "]") {
			viper.Set(k, getVersionOrPanic())
		}
	}

	err = viper.Unmarshal(c)
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}

func getVersionOrPanic() string {
	var version string

	file, err := os.Open("version.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read line by line
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			version = scanner.Text()
			break
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return version
}
