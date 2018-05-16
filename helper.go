package skiplist

type Debugger struct {
	list *SkipList
}

func NewDebugger(l *SkipList) *Debugger {
	return &Debugger{l}
}

func (d *Debugger) Scores() []Score {
	return d.LevelScores(0)
}

func (d *Debugger) Levels() int {
	return len(d.list.root.next)
}

func (d *Debugger) LevelScores(lv int) []Score {
	scores := []Score{}
	if lv >= 0 && lv < d.Levels() {
		for e := d.list.root.next[lv]; e != d.list.root; e = e.next[lv] {
			scores = append(scores, e.score)
		}
	}
	return scores
}

type SkipListExt struct {
	*SkipList
}

func NewExt(l *SkipList) *SkipListExt {
	return &SkipListExt{l}
}

func (l *SkipListExt) RemoveRange(from, to Score) {
	list := l.GetRange(from, to)
	for _, e := range list {
		e.Remove()
	}
}

func (l *SkipListExt) RemoveAll(score Score) {
	l.RemoveRange(score, score)
}

func (l *SkipListExt) GetRange(from, to Score) []*Element {
	list := []*Element{}
	l.RangeEach(from, to, func(e *Element) bool {
		list = append(list, e)
		return true
	})
	return list
}

func (l *SkipListExt) GetAll(score Score) []*Element {
	return l.GetRange(score, score)
}

func (l *SkipListExt) Count(score Score) int {
	cnt := 0
	l.RangeEach(score, score, func(e *Element) bool {
		cnt++
		return true
	})
	return cnt
}
