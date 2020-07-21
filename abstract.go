package lig

import (
	"fmt"
	"lig/writer"
)

// a abstract logger which implements most function of ligo.Logger.
//
// It is the caller's responsibility to implements the following method:
//
// 1. hook func(level Level, msg string, a ...interface{}) string
//
//		a hook function before call write. could be nil
//
// 2. write func(string)
//
//		handler the final string
type abstractLogger struct {
	level Level
	hook  WriteHookFunc
	write writer.WriteFunc
}

type WriteHookFunc func(level Level, msg string) string

func (l *abstractLogger) Level() Level {
	return l.level
}

func (l *abstractLogger) DebugAble() bool {
	return l.level <= DEBUG
}

func (l *abstractLogger) InfoAble() bool {
	return l.level <= INFO
}

func (l *abstractLogger) WarnAble() bool {
	return l.level <= WARN
}

func (l *abstractLogger) ErrorAble() bool {
	return l.level <= ERROR
}

func (l *abstractLogger) Debug(msg string, a ...interface{}) {
	if l.DebugAble() {
		l.beforeWrite(DEBUG, msg, a...)
	}
}

func (l *abstractLogger) Info(msg string, a ...interface{}) {
	if l.InfoAble() {
		l.beforeWrite(INFO, msg, a...)
	}
}

func (l *abstractLogger) Warn(msg string, a ...interface{}) {
	if l.WarnAble() {
		l.beforeWrite(WARN, msg, a...)
	}
}

func (l *abstractLogger) Error(msg string, a ...interface{}) {
	if l.ErrorAble() {
		l.beforeWrite(ERROR, msg, a...)
	}
}

func (l *abstractLogger) beforeWrite(level Level, msg string, a ...interface{}) {
	messageText := msg
	if len(a) > 0 {
		messageText = fmt.Sprintf(msg, a...)
	}
	if l.hook != nil {
		l.write(l.hook(level, messageText))
	} else {
		l.write(messageText)
	}
}
