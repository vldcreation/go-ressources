package util

import (
	"path/filepath"
	"runtime"
)

func RootPath() string {
	return Path(0, 1)
}

func Path(skip, up int) string {
	_, path, _, _ := runtime.Caller(skip)

	for i := 0; i < up; i++ {
		path = filepath.Dir(path)
	}

	return filepath.Dir(path)
}
