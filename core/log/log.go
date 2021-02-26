//Author xc, Created on 2020-04-09 10:00
//{COPYRIGHTS}
package log

import (
	"context"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type ContextInfo struct {
	DebugID   string
	IP        string
	RequestID string
	UserID    int
	Timers    map[string]TimePoint
}

type TimePoint struct {
	Start int64
	End   int64
}

//system info
func Info(message interface{}) {
	log.Info(message)
}

func Warning(message interface{}, label string, ctx ...context.Context) {
	if len(ctx) == 1 {
		fields := GetContextFields(ctx[0])
		log.WithFields(fields).Warning(message, label)
	} else {
		log.Warning(message)
	}
}

//Write error
func Error(message interface{}, label string, ctx ...context.Context) {
	if len(ctx) == 1 {
		fields := GetContextFields(ctx[0])
		log.WithFields(fields).Error(message, label)
	} else {
		log.Error(message, label)
	}
}

func Fatal(message interface{}) {
	log.Fatal(message)
}

//Output debug info with on category.
func Debug(message interface{}, category string, ctx ...context.Context) {
	if len(ctx) == 1 {
		fields := GetContextFields(ctx[0])
		log.WithFields(fields).Debug(message, "["+category+"]")
	} else {
		log.Debug(message, "["+category+"]")
	}
}

func GetContextFields(ctx context.Context) log.Fields {
	info := GetContextInfo(ctx)
	fields := log.Fields{}
	fields["ip"] = info.IP
	fields["request_id"] = info.RequestID
	fields["user_id"] = info.UserID
	return fields
}

type logKey struct{}

// init a context log
func InitContext(ctx context.Context, info *ContextInfo) context.Context {
	newContext := context.WithValue(ctx, logKey{}, info)
	return newContext
}

func GetContextInfo(ctx context.Context) *ContextInfo {
	return ctx.Value(logKey{}).(*ContextInfo)
}

//start timing
func StartTiming(ctx context.Context, category string) {
	info := GetContextInfo(ctx)
	now := time.Now().UnixNano()
	timer := TimePoint{Start: now}
	if info.Timers == nil {
		info.Timers = map[string]TimePoint{}
	}

	info.Timers[category] = timer
}

//End timing on a category
func EndTiming(ctx context.Context, category string) {
	info := GetContextInfo(ctx)

	now := time.Now().UnixNano()
	timer := info.Timers[category]
	timer.End = now

	info.Timers[category] = timer
}

//Log all timing, usually done in the end of request
func LogTiming(ctx context.Context) {
	info := GetContextInfo(ctx)
	for category, timer := range info.Timers {
		duration := int((timer.End - timer.Start) / 1000000)
		Debug(strconv.Itoa(duration)+"ms", category, ctx)
	}
}

func init() {
	//todo: log it to file based on parameters
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.0000"})
}
