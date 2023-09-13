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

	return lc
}

// clone 深度拷贝 zapLogger（用户打上烙印，服务器白纸一张，可不能混淆；新建拷贝一份）
func (l *zapLoggerImpl) clone() *zapLoggerImpl {
	lc := *l
	return &lc
}
