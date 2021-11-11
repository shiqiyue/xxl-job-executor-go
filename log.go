package xxl

import (
	"context"
	"fmt"
	"log"
)

//应用日志
type LogFunc func(req LogReq, res *LogRes) []byte

//系统日志
type Logger interface {
	Info(ctx context.Context, format string, a ...interface{})
	Error(ctx context.Context, format string, a ...interface{})
}

type logger struct {
}

func (l *logger) Info(ctx context.Context, format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(format, a...))
}

func (l *logger) Error(ctx context.Context, format string, a ...interface{}) {
	log.Println(fmt.Sprintf(format, a...))
}
