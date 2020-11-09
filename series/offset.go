package series

import "github.com/dawnkosmos/fastpine/exchange"

type oFFSET struct {
	src    Series
	offset int
	ug     *exchange.UpdateGroup
}

/*
Offset shifts a series src offset bars to the right
Pine script code would be: src[1]
Equivalent here: Offset(src, 1)
*/
func Offset(src Series, offset int) Series {
	return &oFFSET{src, offset, src.UpdateGroup()}
}

func (o *oFFSET) Starttime() int64 {
	return o.src.Starttime() + int64(o.src.Resolution()*(o.offset))
}

func (o *oFFSET) Resolution() int {
	return o.src.Resolution()
}

func (o *oFFSET) Value(index int) float64 {
	return o.src.Value(index + o.offset)
}

func (o *oFFSET) Data() []float64 {
	f := o.src.Data()
	out := f[:len(f)-o.offset]
	return out
}

func (o *oFFSET) UpdateGroup() *exchange.UpdateGroup { return o.ug }
