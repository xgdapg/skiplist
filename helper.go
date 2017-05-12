package skiplist

type orderType struct {
	Asc  CompareFunc
	Desc CompareFunc
}

type orderBy struct {
	Int     orderType
	Int64   orderType
	Float64 orderType
	String  orderType
}

var OrderBy = orderBy{
	Int: orderType{
		Asc: func(lhs, rhs interface{}) bool {
			return lhs.(int) < rhs.(int)
		},
		Desc: func(lhs, rhs interface{}) bool {
			return lhs.(int) > rhs.(int)
		},
	},
	Int64: orderType{
		Asc: func(lhs, rhs interface{}) bool {
			return lhs.(int64) < rhs.(int64)
		},
		Desc: func(lhs, rhs interface{}) bool {
			return lhs.(int64) > rhs.(int64)
		},
	},
	Float64: orderType{
		Asc: func(lhs, rhs interface{}) bool {
			return lhs.(float64) < rhs.(float64)
		},
		Desc: func(lhs, rhs interface{}) bool {
			return lhs.(float64) > rhs.(float64)
		},
	},
	String: orderType{
		Asc: func(lhs, rhs interface{}) bool {
			return lhs.(string) < rhs.(string)
		},
		Desc: func(lhs, rhs interface{}) bool {
			return lhs.(string) > rhs.(string)
		},
	},
}

type Debugger struct {
	list *SkipList
}

func NewDebugger(l *SkipList) *Debugger {
	return &Debugger{l}
}

func (d *Debugger) Keys() []interface{} {
	return d.LevelKeys(0)
}

func (d *Debugger) Levels() int {
	return len(d.list.root.next)
}

func (d *Debugger) LevelKeys(lv int) []interface{} {
	keys := []interface{}{}
	if lv >= 0 && lv < d.Levels() {
		for e := d.list.root.next[lv]; e != d.list.root; e = e.next[lv] {
			keys = append(keys, e.key)
		}
	}
	return keys
}

type SkipListExt struct {
	*SkipList
}

func NewExt(l *SkipList) *SkipListExt {
	return &SkipListExt{l}
}

func (l *SkipListExt) RemoveRange(from, to interface{}) {
	list := l.GetRange(from, to)
	for _, e := range list {
		e.Remove()
	}
}

func (l *SkipListExt) RemoveAll(key interface{}) {
	l.RemoveRange(key, key)
}

func (l *SkipListExt) GetRange(from, to interface{}) []*Element {
	list := []*Element{}
	l.RangeEach(from, to, func(e *Element) bool {
		list = append(list, e)
		return true
	})
	return list
}

func (l *SkipListExt) GetAll(key interface{}) []*Element {
	return l.GetRange(key, key)
}

func (l *SkipListExt) Count(key interface{}) int {
	cnt := 0
	l.RangeEach(key, key, func(e *Element) bool {
		cnt++
		return true
	})
	return cnt
}
