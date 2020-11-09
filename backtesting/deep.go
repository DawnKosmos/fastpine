package backtesting

import (
	"github.com/dawnkosmos/fastpine/series"
)

type deep struct {
	o *series.OHCLV

	conditions []Cons
}
