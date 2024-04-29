package main

import (
	"log/slog"
	"os"

	"github.com/callegarimattia/json_parser/internal/parser"
	"github.com/lmittmann/tint"
)

func main() {
	l := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(l)

	if len(os.Args) < 2 {
		slog.Error("Too few arguments")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		slog.Error("Failed to read the file", "error", err)
		os.Exit(1)
	}

	slog.Debug("Read", "data", string(data))

	p := parser.NewParser()
	code, err := p.Parse(data)
	if err != nil {
		slog.Error("Failed to parse the data", "error", err)
		os.Exit(code)
	}

	slog.Info("JSON parsed", "code", code)

	os.Exit(code)
}
