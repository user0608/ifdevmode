# ifdevmode

A small Go package to check whether your application is running in development mode, and optionally execute a block of code only when development mode is enabled.

## Installation

```bash
go get github.com/user0608/ifdevmode
```

Replace `github.com/user0608/ifdevmode` with your real module path.

## Usage

### Execute code only in development mode

```go
package main

import (
	"context"
	"log"

	"github.com/user0608/ifdevmode"
)

func main() {
	ifdevmode.Do(func() {
		err := connection.Conn(context.Background()).Schema.Create(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
	})
}
```

By default, `Do` runs the function in a goroutine when development mode is enabled.

### Execute synchronously

```go
ifdevmode.Do(func() {
	// development-only code
}, ifdevmode.WithSyncExecution())
```

### Check development mode directly

```go
if ifdevmode.Yes() {
	// development-only code
}
```

### Use a custom condition

```go
ifdevmode.Do(func() {
	// code executed only if the custom condition returns true
}, ifdevmode.WithExecuteOn(func() bool {
	return true
}))
```

## Environment variables

`ifdevmode` supports two types of environment variables:

1. Boolean development flags
2. Application environment mode variables

## Boolean development flags

These variables are interpreted as boolean flags:

```bash
DEV
DEV_MODE
DEVMODE
APP_DEV
APP_DEV_MODE
```

Accepted values:

```bash
true
1
on
enabled
enable
active
activo
si
sí
yes
y
```

Examples:

```bash
DEV=true
DEV_MODE=1
APP_DEV=enabled
APP_DEV_MODE=yes
```

Any of those values will enable development mode.

## Environment mode variables

These variables are interpreted as application mode values:

```bash
ENV
ENVIRONMENT
APP_ENV
APP_MODE
GO_ENV
NODE_ENV
```

Accepted development values:

```bash
dev
develop
development
local
localhost
```

Examples:

```bash
ENV=development
ENVIRONMENT=dev
APP_ENV=local
NODE_ENV=development
```

Any of those values will enable development mode.

## Important behavior

Mode variables such as `ENV`, `ENVIRONMENT`, `APP_ENV`, `APP_MODE`, `GO_ENV`, and `NODE_ENV` do not accept boolean values.

For example, this will not enable development mode:

```bash
ENVIRONMENT=true
```

Use this instead:

```bash
ENVIRONMENT=development
```

Boolean values are only accepted by boolean development flags such as `DEV`, `DEV_MODE`, `DEVMODE`, `APP_DEV`, and `APP_DEV_MODE`.

## Examples

### Enabled

```bash
DEV=true
```

```go
ifdevmode.Yes() // true
```

### Enabled through application mode

```bash
APP_ENV=development
```

```go
ifdevmode.Yes() // true
```

### Disabled in production

```bash
APP_ENV=production
```

```go
ifdevmode.Yes() // false
```

### Disabled with boolean false

```bash
DEV=false
```

```go
ifdevmode.Yes() // false
```

## API

### Yes

```go
func Yes() bool
```

Returns `true` when development mode is detected from the supported environment variables.

### Do

```go
func Do(fn func(), opts ...Option)
```

Executes `fn` only when development mode is enabled.

By default, `fn` is executed asynchronously in a goroutine.

### WithSyncExecution

```go
func WithSyncExecution() Option
```

Makes `Do` execute the function synchronously.

### WithExecuteOn

```go
func WithExecuteOn(fn func() bool) Option
```

Overrides the condition used by `Do`.

## Notes

This package is intentionally focused on development mode detection. Debug-specific environment variables such as `DEBUG`, `DEBUG_MODE`, and framework-specific values such as `GIN_MODE=debug` are not treated as development mode by default.