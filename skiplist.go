package skiplist

import (
	"math/rand"
	"time"
)

var MaxLevel = 32

type Score interface {
	EqualTo(interface{}) bool
	LessThan(interface{}) bool
}

type SkipList struct {
	root *Element
	len  int
}

func New() *SkipList {
	l := &SkipList{
		root: &Element{},
		len:  0,
	}
	l.root.list = l
	l.root.prev = []*Element{l.root}
	l.root.next = []*Element{l.root}
	return l
}

func (l *SkipList) Len() int {
	return l.len
}

func (l *SkipList) Front() *Element {
	return l.root.Next()
}

func (l *SkipList) Back() *Element {
	return l.root.Prev()
}

func (l *SkipList) search(score Score, fastreturn bool) (*Element, []*Element) {
	lv := len(l.root.next) - 1
	path := make([]*Element, lv+1)
	e := l.root.next[lv]
	for lv >= 0 {
		if e != l.root {
			if (fastreturn || lv == 0) && score.EqualTo(e.score) {
				for i := lv; i >= 0; i-- {
					path[i] = e
				}
				return e, path
			}
			if e.score.LessThan(score) {
				e = e.next[lv]
				continue
			}
		}
		path[lv] = e
		if lv == 0 {
			break
		}
		e = e.prev[lv]
		lv--
		e = e.next[lv]
	}
	return nil, path
}

func (l *SkipList) revsearch(score Score, fastreturn bool) (*Element, []*Element) {
	lv := len(l.root.prev) - 1
	path := make([]*Element, lv+1)
	e := l.root.prev[lv]
	for lv >= 0 {
		if e != l.root {
			if (fastreturn || lv == 0) && score.EqualTo(e.score) {
				for i := lv; i >= 0; i-- {
					path[i] = e
				}
				return e, path
			}
			if score.LessThan(e.score) {
				e = e.prev[lv]
				continue
			}
		}
		path[lv] = e
		if lv == 0 {
			break
		}
		e = e.next[lv]
		lv--
		e = e.prev[lv]
	}
	return nil, path
}

func (l *SkipList) Get(score Score) *Element {
	e, _ := l.search(score, true)
	return e
}

func (l *SkipList) GetFirst(score Score) *Element {
	e, _ := l.search(score, false)
	return e
}

func (l *SkipList) GetLast(score Score) *Element {
	e, _ := l.revsearch(score, false)
	return e
}

func (l *SkipList) RangeEach(from, to Score, f func(*Element) bool) {
	le, lpath := l.search(from, false)
	if le == nil {
		le = lpath[0]
		if le == l.root {
			return
		}
	}

	re, rpath := l.revsearch(to, false)
	if re == nil {
		re = rpath[0]
		if re == l.root {
			return
		}
	}

	if !le.score.EqualTo(re.score) && !le.score.LessThan(re.score) {
		return
	}

	for e := le; e != nil; e = e.Next() {
		if !f(e) || e == re {
			break
		}
	}
	return
}

func (l *SkipList) insert(score Score, value interface{}, path []*Element) *Element {
	e := &Element{
		list:  l,
		score: score,
		Value: value,
		prev:  []*Element{},
		next:  []*Element{},
	}

	rndLv := randomLevel()
	for lv := 0; lv <= rndLv; lv++ {
		if lv < len(path) {
			prev := path[lv]
			e.next = append(e.next, prev.next[lv])
			e.prev = append(e.prev, prev)
			e.next[lv].prev[lv] = e
			e.prev[lv].next[lv] = e
		} else {
			e.prev = append(e.prev, l.root)
			e.next = append(e.next, l.root)
			l.root.prev = append(l.root.prev, e)
			l.root.next = append(l.root.next, e)
			break
		}
	}

	return e
}

func (l *SkipList) Add(score Score, value interface{}) *Element {
	_, path := l.revsearch(score, false)
	return l.insert(score, value, path)
}

func (l *SkipList) Set(score Score, value interface{}) *Element {
	e, path := l.revsearch(score, true)
	if e != nil {
		e.Value = value
		return e
	}
	return l.insert(score, value, path)
}

func (l *SkipList) Remove(score Score) *Element {
	e := l.Get(score)
	if e != nil {
		e.Remove()
	}
	return e
}

//

var P int64 = 4

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomLevel() int {
	lv := 0
	n := rnd.Int63()
	i := P
	for lv < MaxLevel && n%i == 0 {
		lv++
		i *= P
	}
	return lv
}

//

type Element struct {
	score      Score
	Value      interface{}
	prev, next []*Element
	list       *SkipList
}

func (e *Element) Score() Score {
	return e.score
}

func (e *Element) Next() *Element {
	if e.list == nil {
		return nil
	}
	if e.next == nil || len(e.next) == 0 {
		return nil
	}
	n := e.next[0]
	if n == e.list.root {
		return nil
	}
	return n
}

func (e *Element) Prev() *Element {
	if e.list == nil {
		return nil
	}
	if e.prev == nil || len(e.prev) == 0 {
		return nil
	}
	p := e.prev[0]
	if p == e.list.root {
		return nil
	}
	return p
}

func (e *Element) Remove() {
	for lv := 0; lv < len(e.next); lv++ {
		e.next[lv].prev[lv] = e.prev[lv]
		e.prev[lv].next[lv] = e.next[lv]
		e.next[lv] = nil
		e.prev[lv] = nil
	}
	l := e.list
	e.list = nil
	for lv := len(l.root.next) - 1; lv > 0; lv-- {
		if l.root.next[lv] == l.root {
			l.root.next = l.root.next[:lv]
			l.root.prev = l.root.prev[:lv]
		}
		break
	}
}
