package lig

import "fmt"

type prefixLogger struct {
	logger Logger
	prefix string
}

// a logger with a fixed prefix
func PrefixLogger(prefix string, logger Logger) Logger {
	return &prefixLogger{
		logger: logger,
		prefix: prefix,
	}
}

func (l *prefixLogger) DebugAble() bool {
	return l.logger.DebugAble()
}

func (l *prefixLogger) InfoAble() bool {
	return l.logger.InfoAble()
}

func (l *prefixLogger) WarnAble() bool {
	return l.logger.WarnAble()
}

func (l *prefixLogger) ErrorAble() bool {
	return l.logger.ErrorAble()
}

func (l *prefixLogger) Level() Level {
	return l.logger.Level()
}

func (l *prefixLogger) Debug(msg string, a ...interface{}) {
	if l.DebugAble() {
		l.logger.Debug(l.msgWithPrefix(msg, a))
	}
}

func (l *prefixLogger) Info(msg string, a ...interface{}) {
	if l.InfoAble() {
		l.logger.Info(l.msgWithPrefix(msg, a))
	}
}

func (l *prefixLogger) Warn(msg string, a ...interface{}) {
	if l.WarnAble() {
		l.logger.Warn(l.msgWithPrefix(msg, a))
	}
}

func (l *prefixLogger) Error(msg string, a ...interface{}) {
	if l.ErrorAble() {
		l.logger.Error(l.msgWithPrefix(msg, a))
	}
}

func (l *prefixLogger) msgWithPrefix(msg string, a []interface{}) string {
	if len(a) == 0 {
		return l.prefix + msg
	} else {
		return l.prefix + fmt.Sprintf(msg, a...)
	}
}
