# Disposable Email Domains in Golang

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/disposable)
[![Build Status](https://travis-ci.org/ferhatelmas/disposable.svg?branch=master)](https://travis-ci.org/ferhatelmas/disposable)

> A list of disposable and temporary email address domains in Golang.

There are two lists; namely, `Black` and `White`.

`Black`: These are disposable and generally used to register dummy users to abuse/spam other services.
`White`: These are real email providers but generally treated as disposable wrongly.

## Install

By go tool: `go get github.com/ferhatelmas/disposable`

## Usage

```go
import github.com/ferhatelmas/disposable

disposable.IsBlack("google.com") // false
disposable.IsBlack("10mail.org") // true
disposable.IsBlack("10MAIL.org") // true
```

## RELATED

Actual list is collected @ [martersen/disposable-email-domains](https://github.com/martenson/disposable-email-domains).

## CONTRIBUTING

Feel free to create a PR with additions or removal requests here or above related repo. For sure, don't forget adding explanation of request.

## LICENSE

MIT © [Ferhat Elmas](https://github.com/ferhatelmas)
