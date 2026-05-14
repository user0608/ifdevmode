package ifdevmode

import "testing"

func TestEnabledWithBoolEnvKeys(t *testing.T) {
	tests := []struct {
		name string
		key  string
		val  string
	}{
		{name: "DEV true", key: "DEV", val: "true"},
		{name: "DEV_MODE one", key: "DEV_MODE", val: "1"},
		{name: "DEVMODE on", key: "DEVMODE", val: "on"},
		{name: "APP_DEV enabled", key: "APP_DEV", val: "enabled"},
		{name: "APP_DEV_MODE yes", key: "APP_DEV_MODE", val: "yes"},
		{name: "spanish si", key: "DEV", val: "si"},
		{name: "spanish accent sí", key: "DEV", val: "sí"},
		{name: "uppercase value", key: "DEV", val: "TRUE"},
		{name: "trimmed value", key: "DEV", val: "  true  "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lookup := func(key string) string {
				if key == tt.key {
					return tt.val
				}
				return ""
			}

			if !enabled(lookup) {
				t.Fatal("expected dev mode to be enabled")
			}
		})
	}
}

func TestEnabledWithModeEnvKeys(t *testing.T) {
	tests := []struct {
		name string
		key  string
		val  string
	}{
		{name: "ENV dev", key: "ENV", val: "dev"},
		{name: "ENVIRONMENT develop", key: "ENVIRONMENT", val: "develop"},
		{name: "APP_ENV development", key: "APP_ENV", val: "development"},
		{name: "APP_MODE local", key: "APP_MODE", val: "local"},
		{name: "GO_ENV localhost", key: "GO_ENV", val: "localhost"},
		{name: "NODE_ENV uppercase", key: "NODE_ENV", val: "DEVELOPMENT"},
		{name: "NODE_ENV trimmed", key: "NODE_ENV", val: "  development  "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lookup := func(key string) string {
				if key == tt.key {
					return tt.val
				}
				return ""
			}

			if !enabled(lookup) {
				t.Fatal("expected dev mode to be enabled")
			}
		})
	}
}

func TestEnabledReturnsFalseForNonDevValues(t *testing.T) {
	tests := []struct {
		name string
		key  string
		val  string
	}{
		{name: "empty", key: "DEV", val: ""},
		{name: "DEV false", key: "DEV", val: "false"},
		{name: "DEV zero", key: "DEV", val: "0"},
		{name: "ENV production", key: "ENV", val: "production"},
		{name: "ENV prod", key: "ENV", val: "prod"},
		{name: "ENV true should not count", key: "ENV", val: "true"},
		{name: "ENVIRONMENT true should not count", key: "ENVIRONMENT", val: "true"},
		{name: "unknown value", key: "DEV", val: "banana"},
		{name: "unknown key", key: "DEBUG", val: "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lookup := func(key string) string {
				if key == tt.key {
					return tt.val
				}
				return ""
			}

			if enabled(lookup) {
				t.Fatal("expected dev mode to be disabled")
			}
		})
	}
}

func TestEnabledReturnsTrueIfAnyMatchingEnvExists(t *testing.T) {
	lookup := func(key string) string {
		switch key {
		case "ENV":
			return "production"
		case "APP_ENV":
			return "development"
		default:
			return ""
		}
	}

	if !enabled(lookup) {
		t.Fatal("expected dev mode to be enabled")
	}
}

func TestEnabledWithNilLookup(t *testing.T) {
	t.Setenv("DEV", "true")

	if !enabled(nil) {
		t.Fatal("expected dev mode to be enabled")
	}
}

func TestYes(t *testing.T) {
	t.Setenv("APP_ENV", "development")

	if !Yes() {
		t.Fatal("expected Yes to return true")
	}
}

func TestDoDoesNothingWhenDevModeIsDisabled(t *testing.T) {
	called := false

	Do(func() {
		called = true
	}, WithExecuteOn(func() bool {
		return false
	}), WithSyncExecution())

	if called {
		t.Fatal("expected function not to be called")
	}
}

func TestDoExecutesSynchronously(t *testing.T) {
	called := false

	Do(func() {
		called = true
	}, WithExecuteOn(func() bool {
		return true
	}), WithSyncExecution())

	if !called {
		t.Fatal("expected function to be called")
	}
}

func TestDoIgnoresNilFunction(t *testing.T) {
	Do(nil, WithExecuteOn(func() bool {
		return true
	}), WithSyncExecution())
}

func TestDoIgnoresNilOption(t *testing.T) {
	called := false

	Do(func() {
		called = true
	}, nil, WithExecuteOn(func() bool {
		return true
	}), WithSyncExecution())

	if !called {
		t.Fatal("expected function to be called")
	}
}

func TestWithExecuteOnNilKeepsDefault(t *testing.T) {
	t.Setenv("DEV", "true")

	called := false

	Do(func() {
		called = true
	}, WithExecuteOn(nil), WithSyncExecution())

	if !called {
		t.Fatal("expected function to be called")
	}
}
