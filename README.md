# httpdoc

**httpdoc** gives you direct access to HTTP documentation straight from
your terminal.

## Usage

    httpdoc <something>

### Examples

    $ httpdoc 200           # => Doc about the status code
    $ httpdoc content-type  # => Doc about the header field

## Install

    go get github.com/bfontaine/httpdoc

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

This tool is a work in progress.

| Resources                    | Support |
|------------------------------|:-------:|
| Standard status codes        | ✔       |
| Standard header fields       | ✔       |
| Standard methods             | ✔       |

## See Also

* [`rfc`][rfc-cli]: read RFCs in your terminal

[rfc-cli]: https://github.com/bfontaine/rfc#rfc
