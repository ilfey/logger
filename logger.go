package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type LoggerMethods interface {
	Log(level Level, args ...interface{})
	Logf(level Level, format string, args ...interface{})
	Logln(level Level, args ...interface{})

	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Traceln(args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
}

type Logger interface {
	LoggerMethods

	GetOut() io.Writer
	SetOut(out io.Writer)

	GetLevel() Level
	SetLevel(level Level)

	SetTextFormatter(func(level Level, format string, args ...interface{}) string)
}

type LoggerImpl struct {
	Out           io.Writer
	Buffer        *bytes.Buffer
	Level         Level
	ExitFunc      func(code int)
	TextFormatter func(level Level, format string, args ...interface{}) string
}

func New() Logger {
	return &LoggerImpl{
		Out:           os.Stderr,
		Buffer:        &bytes.Buffer{},
		Level:         LevelInfo,
		ExitFunc:      os.Exit,
		TextFormatter: TextFormatter,
	}
}

func (l *LoggerImpl) canLog(level Level) bool {
	return getLevelInt(l.Level) <= getLevelInt(level)
}

func (l *LoggerImpl) log(level Level, msg string) {
	if !l.canLog(level) {
		return
	}

	defer func() {
		l.Buffer.Reset()
	}()

	l.Buffer.WriteString(
		l.TextFormatter(level, msg),
	)

	l.Out.Write(l.Buffer.Bytes())

	if level == LevelPanic {
		panic(l)
	}
}

func (l *LoggerImpl) Log(level Level, args ...interface{}) {
	l.log(
		level,
		fmt.Sprint(args...),
	)
}

func (l *LoggerImpl) Logf(level Level, format string, args ...interface{}) {
	l.log(
		level,
		fmt.Sprintf(format, args...),
	)
}

func (l *LoggerImpl) Logln(level Level, args ...interface{}) {
	l.log(
		level,
		fmt.Sprintln(args...),
	)
}

func (l *LoggerImpl) Trace(args ...interface{}) {
	l.Log(LevelTrace, args...)
}

func (l *LoggerImpl) Tracef(format string, args ...interface{}) {
	l.Logf(LevelTrace, format, args...)
}

func (l *LoggerImpl) Traceln(args ...interface{}) {
	l.Logln(LevelTrace, args...)
}

func (l *LoggerImpl) Debug(args ...interface{}) {
	l.Log(LevelDebug, args...)
}

func (l *LoggerImpl) Debugf(format string, args ...interface{}) {
	l.Logf(LevelDebug, format, args...)
}

func (l *LoggerImpl) Debugln(args ...interface{}) {
	l.Logln(LevelDebug, args...)
}

func (l *LoggerImpl) Info(args ...interface{}) {
	l.Log(LevelInfo, args...)
}

func (l *LoggerImpl) Infof(format string, args ...interface{}) {
	l.Logf(LevelInfo, format, args...)
}

func (l *LoggerImpl) Infoln(args ...interface{}) {
	l.Logln(LevelInfo, args...)
}

func (l *LoggerImpl) Warn(args ...interface{}) {
	l.Log(LevelWarn, args...)
}

func (l *LoggerImpl) Warnf(format string, args ...interface{}) {
	l.Logf(LevelWarn, format, args...)
}

func (l *LoggerImpl) Warnln(args ...interface{}) {
	l.Logln(LevelWarn, args...)
}

func (l *LoggerImpl) Error(args ...interface{}) {
	l.Log(LevelError, args...)
}

func (l *LoggerImpl) Errorf(format string, args ...interface{}) {
	l.Logf(LevelError, format, args...)
}

func (l *LoggerImpl) Errorln(args ...interface{}) {
	l.Logln(LevelError, args...)
}

func (l *LoggerImpl) Fatal(args ...interface{}) {
	l.Log(LevelFatal, args...)
	l.ExitFunc(1)
}

func (l *LoggerImpl) Fatalf(format string, args ...interface{}) {
	l.Logf(LevelFatal, format, args...)
	l.ExitFunc(1)
}

func (l *LoggerImpl) Fatalln(args ...interface{}) {
	l.Logln(LevelFatal, args...)
	l.ExitFunc(1)
}

func (l *LoggerImpl) Panic(args ...interface{}) {
	l.Log(LevelPanic, args...)
}

func (l *LoggerImpl) Panicf(format string, args ...interface{}) {
	l.Logf(LevelPanic, format, args...)
}

func (l *LoggerImpl) Panicln(args ...interface{}) {
	l.Logln(LevelPanic, args...)
}

func (l *LoggerImpl) GetOut() io.Writer {
	return l.Out
}

func (l *LoggerImpl) SetOut(out io.Writer) {
	l.Out = out
}

func (l *LoggerImpl) SetTextFormatter(f func(level Level, format string, args ...interface{}) string) {
	l.TextFormatter = f
}

func (l *LoggerImpl) GetLevel() Level {
	return l.Level
}

func (l *LoggerImpl) SetLevel(level Level) {
	l.Level = level
}
