package slog_dev

import (
	"context"
	"encoding/json"
	"github.com/fatih/color"
	"io"
	"log"
	"log/slog"
	"os"
)

type DevHandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

type DevHandler struct {
	opts DevHandlerOptions
	slog.Handler
	l     *log.Logger
	attrs []slog.Attr
}

func (opts DevHandlerOptions) NewDevHandler(out io.Writer) *DevHandler {
	return &DevHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       log.New(out, "", 0),
	}
}

func (h *DevHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(attr slog.Attr) bool {
		fields[attr.Key] = attr.Value.Any()

		return true
	})

	for _, attr := range h.attrs {
		fields[attr.Key] = attr.Value.Any()
	}

	var b []byte
	var err error

	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msgColor := color.CyanString(r.Message)

	h.l.Println(
		timeStr,
		level,
		msgColor,
		color.WhiteString(string(b)))

	return nil
}

func SetupDevSlog() *slog.Logger {
	opts := DevHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewDevHandler(os.Stdout)

	return slog.New(handler)
}
