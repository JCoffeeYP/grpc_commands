package logger

import (
	"context"
	"google.golang.org/appengine/log"
	"log/slog"
	"os"
	"wsedge_grpc/config"
)

type Logger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
}

type AppLogger struct {
	l *slog.Logger
}

func New(cfg *config.Config) Logger {
	logLevelMap := map[string]slog.Level{
		"DEBUG": slog.LevelDebug,
		"INFO":  slog.LevelInfo,
		"WARN":  slog.LevelWarn,
		"ERROR": slog.LevelError,
	}
	level, ok := logLevelMap[cfg.LogLevel]
	if !ok {
		level = slog.LevelInfo
		log.Errorf(context.Background(), "invalid log level: %s, used default INFO level", cfg.LogLevel)
	}
	l := &AppLogger{
		l: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})),
	}
	return l
}

func (s *AppLogger) Error(msg string, args ...any) {
	s.l.Error(msg, args...)
}

func (s *AppLogger) Info(msg string, args ...any) {
	s.l.Info(msg, args...)
}

func (s *AppLogger) Debug(msg string, args ...any) {
	s.l.Debug(msg, args...)
}
