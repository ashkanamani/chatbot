package jsonhelper

import (
	"encoding/json"
	"log/slog"
	"os"
	"reflect"
)

func Encode[T any](t T) []byte {
	bs, err := json.Marshal(t)
	if err != nil {
		slog.Error("could not encode json", "error", err, "type", reflect.TypeOf(t), "t", t)
		os.Exit(1)
	}
	return bs
}

func Decode[T any](bs []byte) T {
	var t T
	err := json.Unmarshal(bs, &t)
	if err != nil {
		slog.Error("could not decode json", "error", err, "type", reflect.TypeOf(t), "t", t)
		os.Exit(1)
	}
	return t
}
