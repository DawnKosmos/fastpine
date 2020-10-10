package exchange

type Updater interface {
	Update()
	Add()
}

type UpdateGroup struct {
	ID        string
	Frequenzy int
	ug        []*Updater
}

func NewUpdateGroup(ID string, frequenzy int) UpdateGroup {
	return UpdateGroup{ID, frequenzy, []*Updater{}}
}

func (u *UpdateGroup) Add(n Updater) {
	u.ug = append(u.ug, &n)
}

func (u *UpdateGroup) Update() {
	for _, i := range u.ug {
		(*i).Update()
	}
}

func (u *UpdateGroup) AddLast() {
	for _, i := range u.ug {
		(*i).Add()
	}
}
