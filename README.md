# httpdoc

**httpdoc** gives you direct access to some HTTP documentation straight from
your terminal.

## Usage

    httpdoc <something>

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
