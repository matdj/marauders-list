package main

import (
	"os"
	"testing"
)

func TestShouldLoadDefaultIfNoEnvVarPresent(t *testing.T) {
	assertEquals(t, tryDefaultEnvVar("NO_ENV_VAR", "expected default value"), "expected default value")
}

func TestShouldLoadFromEnvVarIfPresent(t *testing.T) {
	expected := "env var value"
	os.Setenv("UNLIKELY_FOO_BAR", expected)

	assertEquals(t, tryDefaultEnvVar("UNLIKELY_FOO_BAR", "unexpected"), expected)
}

func TestShouldReturnLocalhostAndPortIfRunModeEnvVarNotProd(t *testing.T) {
	os.Setenv("RUN_MODE", "")
	assertEquals(t, determineRunMode(), "localhost:8080")
}

func TestShouldReturn8080IfRunModeEnvVarIsProd(t *testing.T) {
	os.Setenv("RUN_MODE", "prod")
	assertEquals(t, determineRunMode(), ":8080")
	os.Setenv("RUN_MODE", "")
}

func assertEquals(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("Actual: '" + actual + "' does not equal: '" + expected + "'")
	}
}
