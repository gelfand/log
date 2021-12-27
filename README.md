# Log

## Simple logging with different verbosity levels

The same API as stdlib has logger, coloured records.

`go get github.com/gelfand/log`

```go
txErr := errors.New("only one reader is being allowed")
log.Errorf("Failed to begin database transaction: %v", txErr)
```

### License

BSD 3-Clause License
