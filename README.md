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
	l.Insert(4, 1)
	l.Insert(34, 2)
	l.Insert(7, 3)
	l.Insert(13, 4)
	l.Insert(35, 5)
	l.Insert(2, 6)
	l.Insert(4, 7)
	l.Upsert(4, 8)

	h := skiplist.NewHelper(l)

	fmt.Println(h.Keys())

	fmt.Println(h.Levels())
	for lv := h.Levels() - 1; lv >= 0; lv-- {
		fmt.Println(h.LevelKeys(lv))
	}
}
```

