package backtesting

type operation interface {
	Trades(trades []*Trade) []*Trade
	Description() string
}

type Operations []operation

func (op *Operations) Get(t []*Trades) ([]Trades, op) {
	if len(op) > 0 {
		return nil, nil
	}
	o := (*op)[0]
	op = op[1:]
	return o.Trades, o
}

//Comp saves the position of 2 indicators
type comp struct {
	p1, p2      int
	description string
}

func (c *comp) Trades(t []*Trade) []Trades {
	var t1, t2 []*Trade
	for _, v := range t {
		if greaterThanSplit(c.p1, c.p2, v) {
			t1 = append(t1, v)
		} else {
			t2 = append(t2, v)
		}
	}
	return [][]*Trade{t1, t2}
}

func Comp(greater bool, p1 int, p2 int, s1, s2 string) *comp {
	d := s1 + ">" + s2
	return &Comp{p1: p1, p2: p2, description: d}
}

func greaterThanSplit(p1, p2 int, t Trade) bool {
	if t.EntryCondition.Indicator[p1] > t.EntryCondition.Indicator[p2] {
		return true
	} else {
		return false
	}
}

func greaterThanVSplit(p1 int, p2 float64, t Trade) bool {
	if t.EntryCondition.Indicator[p1] > p2 {
		return true
	} else {
		return false
	}
}

type compV struct {
	greater     bool
	p1          int
	p2          float64
	description string
}

func CompV(greater bool, p1 int, p2 float64, s1, s2 string) *compV {
	var d string
	if greater {
		d = s1 + ">" + s2
	} else {
		d = s1 + "<" + s2
	}
	return &CompV{greater: greater, p1: p1, p2: p2, description: d}
}

func (c *compV) Trades(t []*Trade) []Trades {
	var t1, t2 []*Trade
	for _, v := range t {
		if greaterThanVSplit(c.p1, c.p2, v) {
			t1 = append(t1, v)
		} else {
			t2 = append(t2, v)
		}
	}

	if c.greater {
		return [][]*Trade{t1, t2}
	} else {
		return [][]*Trade{t2, t1}
	}
}

type part struct {
	p1          int
	div         int
	description string
}

func Part(p1 int, div int, s1, s2 string) *part {
	d := s1 + " separated in Parts with size " + s2
	return &part{p1, div, d}
}

func (p *part) Trades(t []*Trade) []Trades {
	tMap := make(map[int]([]*Trade))
	var i int
	for _, v := range t {
		i = int(v.EntryCondition.Indicator[p.p1]) / (p.div)
		arr, ok := tMap[i]
		if !ok {
			tMap[i] = []*Trade{v}
		} else {
			tMap[i] = append(tMap[i], v)
		}
	}
	var tradi []*Trade
	for _, v := range tMap {
		tradi = append(tradi, v)
	}

	return tradi

}
