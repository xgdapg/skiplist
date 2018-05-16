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
	l := skiplist.New()
	d := skiplist.NewDebugger(l)

	ids := []int{73, 84, 2, 23, 79, 50, 12, 89, 23}
	for _, id := range ids {
		l.Add(Id(id), id)
	}

	for i := 0; i < d.Levels(); i++ {
		fmt.Println(d.LevelScores(i))
	}
}

type Id int

func (a Id) EqualTo(v interface{}) bool {
	if b, ok := v.(Id); ok {
		return a == b
	}
	panic("unexpected type")
}

func (a Id) LessThan(v interface{}) bool {
	if b, ok := v.(Id); ok {
		return a < b
	}
	panic("unexpected type")
}
```

