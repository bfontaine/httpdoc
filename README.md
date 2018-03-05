# httpdoc

**httpdoc** gives you direct access to HTTP documentation straight from
your terminal.

[![GoDoc](https://godoc.org/github.com/bfontaine/httpdoc?status.svg)](https://godoc.org/github.com/bfontaine/httpdoc)
[![Build Status](https://travis-ci.org/bfontaine/httpdoc.svg?branch=master)](https://travis-ci.org/bfontaine/httpdoc)

## Usage

    httpdoc <something>

### Examples

    $ httpdoc 200           # => Doc about the status code
    $ httpdoc content-type  # => Doc about the header field

## Install

    go get github.com/bfontaine/httpdoc

`httpdoc` requires Go 1.2+.

### Using Homebrew

If you use [Homebrew](http://brew.sh) or [Linuxbrew](http://linuxbrew.sh/) you
can install `httpdoc` with one command:

```
brew install bfontaine/utils/httpdoc
```

## Go API

```go
import (
    "fmt"
    "github.com/bfontaine/httpdoc/httpdoc"
)

doc := httpdoc.DefaultDoc
code, _ := doc.GetStatusCode("200")
fmt.Println(code.PrettyString())
```

## Support

| Resources                    | Support |
|------------------------------|:-------:|
| Standard status codes        | ✔       |
| Standard header fields       | ✔       |
| Standard methods             | ✔       |

## See Also

* [`rfc`][rfc-cli]: read RFCs in your terminal

[rfc-cli]: https://github.com/bfontaine/rfc#rfc
