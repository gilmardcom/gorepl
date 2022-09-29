# gorepl

GoRepl enables you to quickly create your own simple REPL in Go!

For a simple demp, please check this out! https://github.com/gilmardcom/gorepl-demo

## Install

```bash
go get -u github.com/gilmardcom/gorepl
```

# Example

Check: https://github.com/gilmardcom/gorepl-demo

## Config

```go
gorepl.Config.ConStatus = gorepl.CsDisonnected
// etc.
```

## Adding your own commands to your REPL

```go
	AddCommand("quit", "q", "quits the repl", func(args []string) (string, error) {
		StopRepl()
		return "", nil
	})

	AddCommand("clear", "c", "clears the screen", func(args []string) (string, error) {
		Clear()
		return "", nil
	})
```

## Ready? Go REPL!

```go
GoRepl()
```

## License

The MIT License (MIT) - see [`LICENSE.md`](https://github.com/gilmardcom/gorepl/LICENSE.md) for more details.


