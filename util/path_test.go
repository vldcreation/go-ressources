package util_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/vldcreation/go-ressources/util"
)

func TestRootPath(t *testing.T) {
	testPath := util.RootPath() + "/test_data/util/config.env"

	_, path, _, _ := runtime.Caller(0)
	expected := filepath.Dir(filepath.Dir(path)) + "/test_data/util/config.env"

	if testPath != expected {
		t.Errorf("Expected %s, got %s", expected, testPath)
	}

}
