# Go slog example

This repo contains a set of examples of structured logs with [log/slog](https://pkg.go.dev/log/slog) package.

> This project follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

## Requisites

- [Go](go.dev) 1.23+

## Examples

- [Minimal](./cmd/minimal/main.go) very first minimal slog example.
- [Minimal JSON Hanlder](./cmd/minimal-json-handler/main.go) example.
- [Minimal Text Hanlder](./cmd/minimal-text-handler/main.go) example. With text handler each log record is formatted according to [Logfmt standard](https://betterstack.com/community/guides/logging/logfmt/).
- [Customizing default logger](./cmd/customizing-default-logger/main.go).
- [Add attributes to log records](./cmd/add-attributes/main.go).
- [Customizing log level](./cmd/customizing-level/).

## Running the examples

`go run <example file>`

For example, to run the [minimal](./cmd/minimal/main.go) example:

```bash
$ go run cmd/minimal/main.go
2024/11/04 23:38:53 INFO Info message
2024/11/04 23:38:53 WARN Warning message
2024/11/04 23:38:53 ERROR Error message
```

### Format and check code issues

```bash
$ go fmt ./... && go vet ./...
```

## References

- [Structured Logging with slog](https://go.dev/blog/slog)
- [Logging in Go with Slog: The Ultimate Guide ](https://betterstack.com/community/guides/logging/logging-in-go/)
