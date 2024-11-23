# Go Cobra CLI Example

This is cli example project. The cli is created with https://github.com/spf13/cobra.

## Go Environment

```sh
$ go version
go version go1.23.2 linux/amd64
```

## How to build & test

* build
  ```sh
  $ make build
  ```
* test
  ```sh
  $ make test
  $ make test-v
  ```

## Command Usage

```sh
$ ./cobra-cli
This is cli example using cobra.
          cobra repository is here https://github.com/spf13/cobra

Usage:
  cobra-cli [flags]
  cobra-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  random-joke get joke randomly from server

Flags:
  -h, --help   help for cobra-cli

Use "cobra-cli [command] --help" for more information about a command.
```