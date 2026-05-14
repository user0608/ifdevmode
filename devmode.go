package ifdevmode

import (
	"os"
	"strings"
)

var boolEnvKeys = []string{
	"DEV",
	"DEV_MODE",
	"DEVMODE",

	"APP_DEV",
	"APP_DEV_MODE",
}

var modeEnvKeys = []string{
	"ENV",
	"ENVIRONMENT",
	"APP_ENV",
	"APP_MODE",
	"GO_ENV",
	"NODE_ENV",
}

var boolPositives = map[string]struct{}{
	"true":    {},
	"1":       {},
	"on":      {},
	"enabled": {},
	"enable":  {},
	"active":  {},
	"activo":  {},
	"si":      {},
	"sí":      {},
	"yes":     {},
	"y":       {},
}

var devModes = map[string]struct{}{
	"dev":         {},
	"develop":     {},
	"development": {},
	"local":       {},
	"localhost":   {},
}

type options struct {
	sync      bool
	executeIf func() bool
}

type Option func(*options)

func WithSyncExecution() Option {
	return func(o *options) {
		o.sync = true
	}
}

func WithExecuteOn(fn func() bool) Option {
	return func(o *options) {
		if fn != nil {
			o.executeIf = fn
		}
	}
}

func Yes() bool {
	return enabled(os.Getenv)
}

func Do(fn func(), opts ...Option) {
	if fn == nil {
		return
	}

	cfg := options{
		executeIf: Yes,
	}

	for _, opt := range opts {
		if opt != nil {
			opt(&cfg)
		}
	}

	if !cfg.executeIf() {
		return
	}

	if cfg.sync {
		fn()
		return
	}

	go fn()
}

func enabled(lookup func(string) string) bool {
	if lookup == nil {
		lookup = os.Getenv
	}

	for _, key := range boolEnvKeys {
		value := normalize(lookup(key))
		if _, ok := boolPositives[value]; ok {
			return true
		}
	}

	for _, key := range modeEnvKeys {
		value := normalize(lookup(key))
		if _, ok := devModes[value]; ok {
			return true
		}
	}

	return false
}

func normalize(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}
