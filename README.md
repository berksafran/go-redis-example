## Minimal Example of Using Redis by Go with Cobra CLI

```shell
go run main.go

A minimal example of using Redis by Go

Usage:
  go run main.go [command]

Available Commands:
  help        Help about any command
  normal      Run with Normal Mode including GET, SET, EXISTS examples
  pub         Run with Publisher Mode
  sub         Run with Subscription Mode

Flags:
  -c, --channel_name string   Channel Name (required) (default "default_chan")
  -h, --help                  help for go-redis-example

Use "go run main.go [command] --help" for more information about a command.
```