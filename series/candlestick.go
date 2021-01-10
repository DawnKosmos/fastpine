package series

/*
	0 Down|Down
	1 Down|Up
	2 Up | Down
	3 Up | Up
*/
func Last2Bars(src Series) Series {
	src1, src2 := Offset(src, 1), Offset(src, 2)
	c1 := GreaterEqual(src, src1)
	c2 := GreaterEqual(src1, src2)

	r1 := Iff(c1, 0, 2)
	r2 := Iff(c2, 0, 1)
	return Add(r1, r2)
}


/*TODO

Japanese Candlestickbars

*/
