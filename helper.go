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

type Helper struct {
	list *SkipList
}

func NewHelper(l *SkipList) *Helper {
	return &Helper{l}
}

func (h *Helper) Keys() []interface{} {
	return h.LevelKeys(0)
}

func (h *Helper) Levels() int {
	return len(h.list.root.next)
}

func (h *Helper) LevelKeys(lv int) []interface{} {
	keys := []interface{}{}
	if lv >= 0 && lv < h.Levels() {
		for e := h.list.root.next[lv]; e != h.list.root; e = e.next[lv] {
			keys = append(keys, e.key)
		}
	}
	return keys
}
