package series

import "github.com/dawnkosmos/fastpine/exchange"

type OFFSET struct {
	src    Series
	offset int
	ug     *exchange.UpdateGroup
}

func Offset(src Series, offset int) *OFFSET {
	return &OFFSET{src, offset, src.UpdateGroup()}
}

func (o *OFFSET) Resolution() int {
	return o.src.Resolution()
}

func (o *OFFSET) Starttime() int64 {
	return o.src.Starttime() + int64(o.src.Resolution()*(o.offset))
}

func (o *OFFSET) Value(index int) float64 {
	return o.src.Value(index + o.offset)
}

func (o *OFFSET) Data() *[]float64 {
	f := *o.src.Data()
	out := f[o.offset:]
	return &out
}

func (o *OFFSET) UpdateGroup() *exchange.UpdateGroup {
	return o.ug
}
