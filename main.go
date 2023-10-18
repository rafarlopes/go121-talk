package main

import (
	"cmp"
	"errors"
	"log/slog"
	"maps"
	"os"
	"slices"
)

func main() {
	out := min(1, 2, 4)
	slog.Info("min value", "min", out)

	out2 := max(1, 2, 4)
	slog.Info("values", "min", out, "max", out2)

	arr := []int{1, 6, 3, 5}
	slog.Info("before cleaning", "arr", arr)

	clear(arr)
	slog.Info("after cleaning", "arr", arr)

	dict := map[string]string{
		"key": "value",
	}
	slog.Info("before cleaning", "dict", dict)

	clear(dict)
	slog.Info("after cleaning", "dict", dict)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger = logger.With("important", "super important")
	logger.Info("values", "min", out, "max", out2)

	logger.WithGroup("request").Info("received request", "method", "GET", "url", "/accounts/admin")
	// There's is also a `testing/slogtest` to allow you to test your custom handlers

	// New slices package
	numbers := []int{2, 1, 5, 10, 9}
	logger.Info("slices", "max", slices.Max(numbers))
	logger.Info("slices", "min", slices.Min(numbers))
	logger.Info("before reverse", "data", numbers)
	slices.Reverse(numbers)
	logger.Info("reverse", "data", numbers)

	// New maps pacakge
	old := map[string]string{
		"key1": "value1",
	}
	dst := map[string]string{}
	maps.Copy(dst, old)
	logger.Info("copied map", "map", dst)

	// New cmp pacakge
	logger.Info("compare 1 - 1", "res", cmp.Compare(1, 1))
	logger.Info("compare 1 - 2", "res", cmp.Compare(1, 2))
	logger.Info("compare 2 - 1", "res", cmp.Compare(2, 1))

}

func notSupported() error {
	return errors.ErrUnsupported
}
