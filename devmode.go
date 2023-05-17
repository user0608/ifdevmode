package ifdevmode

import (
	"os"
	"strings"
)

var devmode = []string{
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

func Do(f func()) {
	runFunc := false
	for _, mode := range devmode {
		if mode != "" {
			if _, ok := positives[mode]; ok {
				runFunc = true
				break
			}
		}
	}
	if runFunc {
		go f()
	}
}
