package backtesting

type Tree struct {
	Parent        *Tree
	Child         []*Tree
	Operations    []Operation
	UsedOperation []Operation
	trades        []*Trade
	r             Result
}

func NewTree(Operations []Operation, trades []*Trade) *Tree {
	var UsedOp []op = make([]op, 0, len(Operations))
	//Calc result with trades
	t := newTree(nil, trades, Operations, UsedOp, r)
	return t
}

func newTree(Parent *Tree, trades []*Trade, o Operations, UsedOp []Operation) *Tree {
	var t Tree = Tree{Parent: Parent, Operations: o, UsedOperation: UsedOp, trades: trades}
	if len(trades) < 5 || len(Operations) == 0 {
		t.Child = nil
		//Calc Result
		return t
	}
	tr, UsedOp := o.Get(trades)
	t.UsedOperation = append(t.UsedOperation, UsedOp)
	for _, v := range tr {
		t.Child = append(t.Child, newTree(t, v, o, t.UsedOperation))
	}
	return t
}
