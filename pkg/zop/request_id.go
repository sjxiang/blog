package zop

import (
	"context"
	
	"go.uber.org/zap"
)

// C 解析传入的 contetx，尝试提取关注的键值，并添加到 zap.Logger 结构化日志中
func C(ctx context.Context) *zapLoggerImpl {
	return std.C(ctx)
}

func (l *zapLoggerImpl) C(ctx context.Context) *zapLoggerImpl {
	lc := l.clone()

	if requestID := ctx.Value("X-Request-ID"); requestID != nil {
		lc.z = lc.z.With(zap.Any("X-Request-ID", requestID))
	}
	if userId := ctx.Value("X-Username"); userId != nil {
		lc.z = lc.z.With(zap.Any("X-Username", userId))
	}

	return lc
}

// clone 深度拷贝 zapLogger（新建拷贝一份 logger，携带用户的 uuid，防止混淆；旧的 logger 没影响）
func (l *zapLoggerImpl) clone() *zapLoggerImpl {
	lc := *l
	return &lc
}
