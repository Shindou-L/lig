package lig

import (
	"strconv"
	"strings"
)

// 日志记录器
type Logger interface {
	Debug(msg string, a ...interface{})
	Info(msg string, a ...interface{})
	Warn(msg string, a ...interface{})
	Error(msg string, a ...interface{})
	DebugAble() bool
	InfoAble() bool
	WarnAble() bool
	ErrorAble() bool
	Level() Level
}

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

var levelMap = make(map[string]Level)

func init() {
	for level := DEBUG; level <= ERROR; level++ {
		levelMap[strings.ToUpper(level.ToString())] = level
	}
}

// 根据level字符串来获取对应级别的Level对象
//
// Level对象不存在时,则返回比最高级Level还高的Level(方便处理不想打印任何日志的场合)
func OfLevel(lv string) Level {
	level, ok := levelMap[strings.ToUpper(lv)]
	if ok {
		return level
	}
	return 100
}

func (l Level) ToString() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return strconv.Itoa(int(l))
	}
}
