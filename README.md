## ifdevmode
a simple function, to execute a block of code if an environment variable is enabled.

```go
    ifdevmode.Do(func() {
            var err = connection.Conn(context.Background()).Schema.Create(context.Background())
            if err != nil {
                log.Fatalln()
            }
    })
```

Environment variables to execute
```bash
    DEBUG
    DEBUG_MODE
    DEBUGMODE
    DEV_MODE
    DEVMODE
    ## possible values:
    #   - true
    #   - 1
    #   - on
    #   - enabled
    #   - activo
    #   - si
    #   - yes
    #   - y
```