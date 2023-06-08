package ifdevmode

import (
	"os"
	"strings"
)

var devmode = []string{
	strings.ToLower(strings.TrimSpace(os.Getenv("DEBUG"))),
	strings.ToLower(strings.TrimSpace(os.Getenv("DEBUG_MODE"))),
	strings.ToLower(strings.TrimSpace(os.Getenv("DEBUGMODE"))),
	strings.ToLower(strings.TrimSpace(os.Getenv("DEV_MODE"))),
	strings.ToLower(strings.TrimSpace(os.Getenv("DEVMODE"))),
}
var positives = map[string]struct{}{
	"true":    {},
	"1":       {},
	"on":      {},
	"enabled": {},
	"activo":  {},
	"si":      {},
	"yes":     {},
	"y":       {},
}

type options struct {
	issync    bool
	executeif func() bool
}

type option func(o *options)

func WithSyncExecution() option { return func(o *options) { o.issync = true } }

func WithExecuteOn(fn func() bool) option { return func(o *options) { o.executeif = fn } }

func Yes() bool {
	var isdevmode bool
	for _, mode := range devmode {
		if mode != "" {
			if _, ok := positives[mode]; ok {
				isdevmode = true
				break
			}
		}
	}
	return isdevmode
}

func Do(fn func(), withs ...option) {
	var opts options
	opts.executeif = Yes
	for _, wth := range withs {
		wth(&opts)
	}
	if opts.executeif() {
		if opts.issync {
			fn()
		} else {
			go fn()
		}
	}
}
