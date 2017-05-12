# skiplist

## Installation

    go get github.com/xgdapg/skiplist

## GoDoc
[http://godoc.org/github.com/xgdapg/skiplist](http://godoc.org/github.com/xgdapg/skiplist)

## Example
```go
package main

import (
	"fmt"
	"github.com/xgdapg/skiplist"
)

func main() {
	skiplist.P = 2
	l := skiplist.New(skiplist.OrderBy.Int.Asc)
	l.Set(4, 1)
	l.Get(4)
	l.Add(4, 2)
	l.GetFirst(4)
	l.GetLast(4)
	l.Remove(4)

    d := skiplist.NewDebugger(l.SkipList)
    fmt.Println(d.Keys())
}
```

