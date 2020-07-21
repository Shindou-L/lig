package main

import (
	"lig"
	"lig/hook"
	"lig/writer"
	"log"
	"os"
	"time"
)

func main() {
	console := lig.NewBuilder(writer.Standard).WithLevel(lig.INFO).Build()
	console.Debug("this is %v info", "debug")
	console.Info("this is %v info", "info")

	consoleWithLevel := lig.NewBuilder(writer.Standard).WithBeforeWriteHook(hook.WithLevelPrefix).Build()
	consoleWithLevel.Debug("this message should with level prefix")
	consoleWithLevel.Info("this message should with level prefix")

	prefixLogger := lig.PrefixLogger("[any prefix] ", consoleWithLevel)
	prefixLogger.Debug("today is %v", time.Now())

	f, err := os.OpenFile("lig.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	fileLogger := lig.NewBuilder(writer.WithLogger(log.New(f, "", log.LstdFlags))).
		WithBeforeWriteHook(hook.WithLevelPrefix).
		WithLevel(lig.WARN).
		Build()
	fileLogger.Debug("this message should not print, level: %v", "debug")
	fileLogger.Info("this message should not print, level: %v", "info")
	fileLogger.Warn("warning message")
	fileLogger.Error("error message")
}
