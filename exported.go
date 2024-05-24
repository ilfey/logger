package logger

import (
	"io"
)

var (
	std = New()
)

func Log(level Level, args ...interface{}) {
	std.Log(level, args...)
}

func Logf(level Level, format string, args ...interface{}) {
	std.Logf(level, format, args...)
}

func Logln(level Level, args ...interface{}) {
	std.Logln(level, args...)
}

func Trace(args ...interface{}) {
	std.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	std.Tracef(format, args...)
}

func Traceln(args ...interface{}) {
	std.Traceln(args...)
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

func Debugln(args ...interface{}) {
	std.Debugln(args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	std.Infoln(args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

func Warnln(args ...interface{}) {
	std.Warnln(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	std.Errorln(args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	std.Fatalln(args...)
}

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

func GetOut() io.Writer {
	return std.GetOut()
}

func SetOut(out io.Writer) {
	std.SetOut(out)
}

func SetTextFormatter(f func(level Level, format string, args ...interface{}) string) {
	std.SetTextFormatter(f)
}

func GetLevel() Level {
	return std.GetLevel()
}

func SetLevel(level Level) {
	std.SetLevel(level)
}
