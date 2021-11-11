package xxl

import "context"

type Extension interface {
	// 执行任务前
	Before(ctx context.Context) context.Context

	// 执行任务后
	After(ctx context.Context) context.Context
}
