# Slugify-stdin

## Usage

Send fixed text to STDIN:

```shell
$ go run cmd/slugify.go <<< "foo-bar'\.<?_+413123213"
foo-bar-_-413123213%
```

Send text interactively via STDIN and press CTRL-D to close STDIN:

```shell
$ go run cmd/slugify.go
hello
/';+))
bar
```

result:

```
hello-bar%
```
