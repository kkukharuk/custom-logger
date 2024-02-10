package logger

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	filelog "gopkg.in/natefinch/lumberjack.v2"
)

func Init(cfg Config) (*Logger, error) {
	l := Logger{}
	l.Log = log.New()
	l.Prefix = cfg.Prefix
	switch cfg.Type {
	case "stdout":
		l.Log.SetOutput(os.Stdout)
	case "stderr":
		l.Log.SetOutput(os.Stderr)
	case "file":
		if cfg.LogFile == "" {
			cfg.LogFile = "logs/diy.server.log"
		}
		if cfg.Size == 0 {
			cfg.Size = 20
		}
		if cfg.Backups == 0 {
			cfg.Backups = 100
		}
		if cfg.Age == 0 {
			cfg.Age = 30
		}
		l.Log.SetOutput(&filelog.Logger{
			Filename:   cfg.LogFile,
			MaxSize:    cfg.Size,
			MaxBackups: cfg.Backups,
			MaxAge:     cfg.Age,
			Compress:   cfg.Compress,
		})
	default:
		l.Log.SetOutput(os.Stdout)
	}
	switch cfg.Format {
	case "text":
		l.Log.SetFormatter(&LogFormatter{})
	case "json":
		l.Log.SetFormatter(&log.JSONFormatter{})
	}
	switch cfg.Level {
	case "error":
		l.Log.SetLevel(log.ErrorLevel)
	case "warn":
		l.Log.SetLevel(log.WarnLevel)
	case "info":
		l.Log.SetLevel(log.InfoLevel)
	case "debug":
		l.Log.SetLevel(log.DebugLevel)
	default:
		l.Log.SetLevel(log.ErrorLevel)
	}
	l.Info("Initialize logger", "init")
	if cfg.Format == "text" {
		l.Info("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇⠄⠸⣿⣿⣿⣿⣿⣿⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠄⠄⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠟⠁⠄⠄⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠁⠄⠄⠄⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⠄⠄⠄⠄⠄⣿⠛⠋⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣠⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⠄⠄⠄⠄⠄⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⠄⠄⠄⠄⠄⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣴⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⠄⠄⠄⠄⠄⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢈⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⠄⠄⣴⣶⡄⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠰⣿⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⣤⣤⣭⣯⣤⣿⣿⣿⣷⣶⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣿⣿⣿⣿", "init")
		l.Info("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿", "init")
	}
	return &l, nil
}

func (formatter *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	logMsg := fmt.Sprintf("%s [%s] %s :: %s :: %s\n", entry.Time.Format("2006-01-02T15:04:05.000000"), entry.Level.String(), entry.Data["system"], entry.Data["module"], entry.Message)
	return []byte(logMsg), nil
}

func (l *Logger) makeFields(module string) log.Fields {
	var mod string
	if len(module) != 0 {
		mod = module
	} else {
		mod = "main"
	}
	result := log.Fields{
		"system": l.Prefix,
		"module": mod,
	}
	return result
}

func (l *Logger) Info(msg string, module string) { l.Log.WithFields(l.makeFields(module)).Info(msg) }

func (l *Logger) Warn(msg string, module string) { l.Log.WithFields(l.makeFields(module)).Warn(msg) }

func (l *Logger) Error(msg string, module string) { l.Log.WithFields(l.makeFields(module)).Error(msg) }

func (l *Logger) Debug(msg string, module string) { l.Log.WithFields(l.makeFields(module)).Debug(msg) }

func (l *Logger) Fatal(msg string, module string) { l.Log.WithFields(l.makeFields(module)).Fatal(msg) }
