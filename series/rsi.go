package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/helper"
)

/* rsi CALC
change = change(close)
gain = change >= 0 ? change : 0.0
loss = change < 0 ? (-1) * change : 0.0
avgGain = rma(gain, 14)
avgLoss = rma(loss, 14)
rs = avgGain / avgLoss
rsi = 100 - (100 / (1 + rs))

RMA CALC
The initial value is a sma

alpha = 1 / len
pine_rma(x, y) =>
	sum := alpha * x + (1 - alpha) * nz(sum[1])
*/

type rsi struct {
	src Series
	len int
	USR
	//starttime  int64
	//resolution int
	/*An Cist element will save following
	source
	avgGain
	avgLose
	*/
	data *cist.Cist
	//ug         *exchange.UpdateGroup
	alpha float64

	tempResult float64
}

//Rsi is equivalent to rsi(src, len) in pinescript
func Rsi(src Series, l int) Series {
	var r rsi
	//INIT
	r.len = l
	r.src = src
	r.resolution = src.Resolution()
	r.starttime = src.Starttime() + int64(r.resolution*l)
	if src.UpdateGroup() != nil {
		r.ug = src.UpdateGroup()
		(*r.ug).Add(&r)
	}
	r.data = cist.New()
	f := src.Data()
	gain, loss := initGainLoss(f)
	r.alpha = 1 / float64(l)
	avgGain, avgLoss := helper.FloatAverage(gain[:l]), helper.FloatAverage(loss[:l])
	rs := avgGain / avgLoss
	rsi := make([]float64, 0, len(f))
	//First rsi Value
	rsi = append(rsi, 100-(100/(1+rs)))
	//Iterating and dynamically calculating the rsi
	for i := l; i < len(gain); i++ {
		avgGain = r.alpha*gain[i] + (1-r.alpha)*avgGain
		avgLoss = r.alpha*loss[i] + (1-r.alpha)*avgLoss
		rsi = append(rsi, (100 - (100 / (1 + avgGain/avgLoss))))
	}

	l1, l2 := len(f), len(gain)
	r.data.FillElements(f[l1-l:], gain[l2-l:])
	r.data.InitData(rsi)
	return &r
}

func (r *rsi) Update() {
	v := r.src.Value(0)
	if v == r.tempResult {
		return
	}
	r.tempResult = v
	rma := r.data.First()
	g, l := isGainLose(Change(rma[0], v))
	avgGain := r.alpha*g + (1-r.alpha)*rma[1]
	avgLoss := r.alpha*l + (1-r.alpha)*rma[2]
	rsi := 100 - (100 / (1 + avgGain/avgLoss))
	r.data.Update(rsi, avgGain, avgLoss)
}

func (r *rsi) Add() {
	r.data.Add()
}

func (r *rsi) Value(index int) float64 {
	return (*r.data).Get(index)
}

func (r *rsi) Data() []float64 {
	return (*r.data).GetData()
}

//Gets you the avg loss/gain for the rsi calculation
func initGainLoss(f []float64) (gain []float64, loss []float64) {
	gain = append(gain, 0)
	loss = append(loss, 0)
	var change float64
	for i := 1; i < len(f); i++ {
		change = Change(f[i-1], f[i])
		if change >= 0 {
			gain = append(gain, change)
			loss = append(loss, 0)
		} else {
			gain = append(gain, 0)
			loss = append(loss, -1*change)
		}
	}

	return
}

func isGainLose(c float64) (g float64, l float64) {
	if c >= 0 {
		return c, 0
	} else {
		return 0, -1 * c
	}
}
