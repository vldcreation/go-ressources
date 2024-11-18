package util

import (
	"path/filepath"
	"runtime"
)

func RootPath() string {
	_, path, _, _ := runtime.Caller(1)
	return filepath.Dir(path)
}
