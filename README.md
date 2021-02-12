# clark

<div align="center">
  <a href="https://golang.org/">
    <img
      src="https://img.shields.io/badge/MADE%20WITH-GO-%23EF4041?style=for-the-badge"
      height="30"
    />
  </a>
  <a href="https://pkg.go.dev/github.com/lukasl-dev/clark">
    <img
      src="https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge"
      height="30"
    />
  </a>
  <a href="https://goreportcard.com/report/github.com/lukasl-dev/clark">
    <img
      src="https://goreportcard.com/badge/github.com/lukasl-dev/clark?style=for-the-badge"
      height="30"
    />
  </a>
</div>

<br>

- [clark](#clark)
  - [What is `clark`?](#what-is-clark)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
    - [Install the package](#install-the-package)
    - [Install the CLI](#install-the-cli)
  - [Getting started](#getting-started)
    - [Create a Reader](#create-a-reader)
    - [Create a Lexer](#create-a-lexer)
      - [Available Options](#available-options)
    - [Start the lexing progress](#start-the-lexing-progress)
    - [Collect the lexed Tokens](#collect-the-lexed-tokens)
  - [Usage of the CLI](#usage-of-the-cli)

---

## What is `clark`?

`Clark` is a simple command-lexing package designed for chat systems such as Discord, Teamspeak, or Telegram. It is
therefore intended to facilitate the handling of user input and the parsing processes that follow.

---

## Prerequisites

To use a Go package such as `clark`, you must of course have Go installed on your system.

It is assumed that you have already worked with the Go environment. If this is not the case,
see [this page first](https://golang.org/doc/install).

---

## Installation

### Install the package

To use `clark` as a Go package, you must have it installed on your current system. If this is not the case, run the
command below.

```console
go get -u github.com/lukasl-dev/clark
```

### Install the CLI

Otherwise, if you want to use `clark` inside your terminal, you need to have it installed as well. If this is not the
case, run the command below.

If you have no experience in this section,
read [this](https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies).

```console
go get -u github.com/lukasl-dev/clark/cmd/clark
```

---

## Getting started

### Create a [Reader](https://pkg.go.dev/io#Reader)

As an example, we create a `strings.Reader` here.

```go
reader := strings.NewReader("!help")
```

### Create a [Lexer](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing#Lexer)

```go
lex, err := clark.NewLexer(reader)
```

You can configure the [lexer](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing#Lexer) using the
variadic [options](https://pkg.go.dev/github.com/lukasl-dev/clark#Option).

See [available Options](#available-options).

```go
lex, err := clark.NewLexer(
  reader,
  lexing.Prefixes("!"), lexing.Labels("help", "play", "help me"),
)

// handle error
```

#### Available [Options](https://pkg.go.dev/github.com/lukasl-dev/clark#Option)

| Function                                                                                            |                                              Description |
| :-------------------------------------------------------------------------------------------------- | -------------------------------------------------------: |
| [`lexing.Prefixes`](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing#Prefixes)                 |         Prefixes defines the prefixes used by the lexer. |
| [`lexing.Labels`](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing#Labels)                     |             Labels defines the labels used by the lexer. |
| [`lexing.PrefixIgnoreCase`](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing#PrefixIgnoreCase) |  PrefixIgnoreCase defines the prefixes case sensitivity. |
| [`lexing.LabelIgnoreCase`](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing#LabelIgnoreCase)   |     LabelIgnoreCase defines the labels case sensitivity. |
| [`lexing.SkipPrefix`](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing#SkipPrefix)             | SkipPrefix defines whether the prefix should be skipped. |
| [`lexing.SkipLabel`](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing#SkipLabel)               |   SkipLabel defines whether the label should be skipped. |

### Start the lexing progress

To execute the [Lexer](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing#Lexer)
, `lex.Lex()` is executed as a goroutine. `lex.Lex()` reads the passed reader lexicographically until `io.EOF` and
returns the lexed [tokens](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing/token#Token).

```go
go lex.Lex()
```

### Collect the lexed [Tokens](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing/token#Token)

During reading [tokens](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing/token#Token) are sent to the token
channel. This channel can be accessed via `lex.Chan()`.

```go
for token := range lex.Chan() {
// handle token
}
```

A [Token](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing/token#Token) has
a [Type](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing/token#Type) and an optional value.

See supported [types](https://pkg.go.dev/github.com/lukasl-dev/clark/lexing/token#Type).

---

## Usage of the CLI
