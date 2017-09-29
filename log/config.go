package log

import (
	"encoding/json"
	"log"
)

//Config for the base Log
type Config struct {
	Appenders map[string]*AppenderConfig `json:"appenders"`
	Loggers   map[string]*LoggerConfig   `json:"loggers"`
}

type LoggerConfig struct {
	Name         string
	Appenders    []string `json:"appenders"`
	Level        int      `json:"levels"`
	ShowFileLine bool     `json:"showFileLine"`
	ShortFile    bool     `json:"shortFile"`
}

type AppenderConfig struct {
	Name string
	//Console or file
	Type     string      `json:"type"`
	FilePath string      `json:"filePath"`
	Levels   interface{} `json:"levels"`
	IsDaliy  bool        `json:"isDaily"`
}

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig(config *Config) {
	for name, appConfig := range config.Appenders {
		appConfig.Name = name
		loadAppender(*appConfig)
	}
	for name, logConfig := range config.Loggers {
		logConfig.Name = name
		loadLogger(*logConfig)
	}
}

func LoadConfigJsonStr(configStr string) {
	var conf = &Config{}
	err := json.Unmarshal([]byte(configStr), conf)
	if err != nil {
		log.Fatalln(err)
		return
	}
	LoadConfig(conf)
}

func loadAppender(config AppenderConfig) {
	var tmp Appender
	switch config.Type {
	case AppenderTypeConsole:
		tmp = newConsoleAppender()
	case AppenderTypeFile:
		tmp = newFileAppender(config)
	case AppenderTypeWeb:
		tmp = newWebAppender()
	default:
		log.Fatalln("error log type ", config.Type)
		return
	}
	tmp.SetLevel(config.Levels)
	_appenders[config.Name] = tmp
}

func loadLogger(config LoggerConfig) {
	l := &Logger{
		LoggerConfig: config,
	}
	for _, ap := range config.Appenders {
		if apr, ok := _appenders[ap]; ok {
			l.appenders = append(l.appenders, apr)
		}
	}
	if len(l.appenders) < 1 {
		log.Fatalf("new logger error from config %#v \n", config)
	}
	_loggers[config.Name] = l
	if config.Name == "root" {
		rootLog = l
	}
}
