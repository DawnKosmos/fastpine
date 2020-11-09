package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/helper"
)

type vwma struct {
	src        Series
	vol        Series
	len        int
	starttime  int64
	resolution int
	ug         *exchange.UpdateGroup

	USR
	data               *cist.Cist
	tempResult         float64
	volSum, volXsrcSum float64
}

//Vwma (src,volume,len) is equivalent to vwma(src,len)
func Vwma(src Series, volume Series, l int) Series {
	var s vwma
	s.len = l
	s.src = src
	s.vol = volume
	s.resolution = volume.Resolution()
	s.starttime = src.Starttime() + int64(l*src.Resolution())
	if src.UpdateGroup() != nil {
		s.ug = src.UpdateGroup()
		(*s.ug).Add(&s)
	}
	s.data = cist.New()
	f, vo := src.Data(), volume.Data()
	fOut := make([]float64, 0, len(f))
	vo = vo[len(vo)-len(f):]
	volSum := helper.FloatSum(vo[:l])
	volXsrcSum := helper.FloatSum(opExecute(mult, vo[:l], f[:l]))
	avg := volXsrcSum / volSum
	fOut = append(fOut, avg)

	for i := l; i < len(f); i++ {
		volSum = volSum + vo[i] - vo[i-l]
		volXsrcSum = volXsrcSum + vo[i]*f[i] - vo[i-l]*f[i-l]
		fOut = append(fOut, volXsrcSum/volSum)
	}

	s.data.InitData(fOut)
	calSrc := make([]float64, 0, l)
	calVol := make([]float64, 0, l)

	for i, v := range vo[len(f)-l:] {
		calSrc = append(calSrc, v)
		calVol = append(calVol, f[i])
	}

	s.data.FillElements(calSrc, calVol)

	fuckS := s.data.Last()
	s.volSum = volSum - fuckS[1]
	s.volXsrcSum = volXsrcSum - fuckS[1]*fuckS[0]

	return &s
}

func (s *vwma) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *vwma) Data() []float64 {
	return (*s.data).GetData()
}

func (s *vwma) Add() {
	s.data.Add()
}

func (s *vwma) Update() {
	src := s.src.Value(0)
	vol := s.vol.Value(0)
	if src*vol == s.tempResult {
		return
	}
	s.tempResult = src * vol

	volXsrcSum := s.volXsrcSum + s.tempResult
	volSum := s.volSum + vol

	s.data.Update(volXsrcSum/volSum, src, vol)
}
