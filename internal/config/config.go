package config

import (
	"fmt"
	"jpnovelmtlgo/internal/exception"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Config interface {
	Get(key string) string
	App() Application
}

var (
	AppConfig *configImpl
)

type configImpl struct {
	Application Application `yaml:"application"`
}

type Application struct {
	AppName        string    `yaml:"appName"`
	AppVersion     string    `yaml:"appVersion"`
	AppDescription string    `yaml:"appDescription"`
	Server         Server    `yaml:"server"`
	Translate      Translate `yaml:"translate"`
}

type Server struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Translate struct {
	Url string `yaml:"url"`
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func (config *configImpl) App() Application {
	return config.Application
}

func New(file string) Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file found")
	}

	data, err := os.ReadFile(file)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	data = []byte(os.ExpandEnv(string(data)))

	err = yaml.Unmarshal(data, &AppConfig)
	exception.PanicIfNeeded(err)

	return AppConfig
}
