package logrgorm2

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormLogger interface {
	logger.Interface
	IgnoreRecordNotFoundError(bool) GormLogger
	SlowThreshold(time.Duration) GormLogger
}

type Logger struct {
	ll                        logr.Logger
	lv                        logger.LogLevel
	slowThreshold             time.Duration
	ignoreRecordNotFoundError bool
}

func New(ll logr.Logger) GormLogger {
	return &Logger{
		ll:                        ll.WithCallDepth(1),
		slowThreshold:             100 * time.Millisecond,
		ignoreRecordNotFoundError: false,
	}
}

func (l *Logger) IgnoreRecordNotFoundError(ignore bool) GormLogger {
	return &Logger{
		ll:                        l.ll,
		lv:                        l.lv,
		slowThreshold:             l.slowThreshold,
		ignoreRecordNotFoundError: ignore,
	}
}

func (l *Logger) SlowThreshold(threshold time.Duration) GormLogger {
	return &Logger{
		ll:                        l.ll,
		lv:                        l.lv,
		slowThreshold:             threshold,
		ignoreRecordNotFoundError: l.ignoreRecordNotFoundError,
	}
}

func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	return &Logger{
		ll:                        l.ll,
		lv:                        level,
		slowThreshold:             l.slowThreshold,
		ignoreRecordNotFoundError: l.ignoreRecordNotFoundError,
	}
}

func (l *Logger) Info(ctx context.Context, s string, i ...interface{}) {
	if l.lv < logger.Info {
		return
	}
	l.ll.V(1).Info(fmt.Sprintf(s, i...))
}

func (l *Logger) Warn(ctx context.Context, s string, i ...interface{}) {
	if l.lv < logger.Warn {
		return
	}
	l.ll.V(0).Info(fmt.Sprintf(s, i...))
}

func (l *Logger) Error(ctx context.Context, s string, i ...interface{}) {
	if l.lv < logger.Error {
		return
	}
	l.ll.Error(fmt.Errorf(s, i...), "gorm error")
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.lv <= 0 {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.lv >= logger.Error && (!l.ignoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		l.ll.Error(err, "Gorm SQL trace", "elapsed", elapsed, "rows", rows, "sql", sql)
	case l.slowThreshold != 0 && elapsed > l.slowThreshold && l.lv >= logger.Warn:
		sql, rows := fc()
		l.ll.V(0).Info("Gorm SQL trace", "elapsed", elapsed, "rows", rows, "sql", sql)
	case l.lv >= logger.Info:
		sql, rows := fc()
		l.ll.V(1).Info("Gorm SQL trace", "elapsed", elapsed, "rows", rows, "sql", sql)
	}
}
