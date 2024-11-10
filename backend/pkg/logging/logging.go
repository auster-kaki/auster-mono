package logging

import (
	"context"
	"log/slog"
	"os"
)

func init() {
	w := os.Stdout
	slog.SetDefault(slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.MessageKey {
				return slog.Attr{Key: "message", Value: a.Value}
			}
			return a
		},
	})))
}

func Info(ctx context.Context, msg string, args ...any) {
	slog.InfoContext(ctx, msg, args)
}

func Warn(ctx context.Context, msg string, args ...any) {
	slog.WarnContext(ctx, msg, args)
}

func Error(ctx context.Context, msg string, args ...any) {
	slog.ErrorContext(ctx, msg, args)
}
