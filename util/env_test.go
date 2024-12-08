package util_test

import (
	"os"
	"testing"

	"github.com/vldcration/go-ressources/util"
)

type TestSuite struct {
	EnvSuite *EnvSuite
}

type EnvSuite struct {
	Foo string `mapstructure:"foo" json:"foo"`
	Bar string `mapstructure:"bar" json:"bar"`
}

func TestLoadEnvFromPath(t *testing.T) {
	t.Run("Test LoadEnvFromPath", func(t *testing.T) {
		expected := &EnvSuite{
			Foo: "foo",
			Bar: "bar",
		}

		actual := &EnvSuite{}
		err := util.LoadEnvFromPath("test_data/util/config.env", actual)
		if err != nil {
			t.Errorf("Error loading env from path: %v", err)
		}

		if actual.Foo != expected.Foo {
			t.Errorf("Expected %s, got %s", expected.Foo, actual.Foo)
		}

		if actual.Bar != expected.Bar {
			t.Errorf("Expected %s, got %s", expected.Bar, actual.Bar)
		}

	})
}

func TestLoadEnvFromBytes(t *testing.T) {
	stream, err := os.ReadFile(util.RootPath() + "/test_data/util/config.json")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	t.Run("Test LoadEnvFromPath", func(t *testing.T) {
		expected := &EnvSuite{
			Foo: "foo",
			Bar: "bar",
		}

		actual := &EnvSuite{}
		err := util.LoadEnvFromBytes(stream, actual)
		if err != nil {
			t.Errorf("Error loading env from bytes: %v", err)
		}

		if actual.Foo != expected.Foo {
			t.Errorf("Expected %s, got %s", expected.Foo, actual.Foo)
		}

		if actual.Bar != expected.Bar {
			t.Errorf("Expected %s, got %s", expected.Bar, actual.Bar)
		}

	})
}
