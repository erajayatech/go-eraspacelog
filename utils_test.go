package eraspacelog

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	resultWithDefaultValue := GetEnv("MY_KEY", "MY_DEFAULT_VALUE")
	if resultWithDefaultValue != "MY_DEFAULT_VALUE" {
		t.Error("Error: default value should be MY_DEFAULT_VALUE")
	}

	os.Setenv("MY_KEY", "sample")
	resultWithoutDefaultValue := GetEnv("MY_KEY", "")
	if resultWithoutDefaultValue == "" {
		t.Error("Error: MY_KEY is empty value")
	}

	t.Log("end of TestGetEnv")
}
