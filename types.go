package logger

import (
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Type     string `json:"type" yaml:"type"`
	Format   string `json:"format" yaml:"format"`
	Level    string `json:"level" yaml:"level"`
	LogFile  string `json:"log_file" yaml:"log_file"`
	Size     int    `json:"size" yaml:"size"`
	Backups  int    `json:"backups" yaml:"backups"`
	Age      int    `json:"age" yaml:"age"`
	Compress bool   `json:"compress" yaml:"compress"`
	Prefix   string `json:"prefix" yaml:"prefix"`
}

type Logger struct {
	Log    *log.Logger
	Prefix string
}

type LogFormatter struct {
}
