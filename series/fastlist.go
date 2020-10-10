package series

//Element saves the previous and next element and a Value(interface{})
type Element struct {
	next, prev *Element
	value      float64
}

/*Fist  is a strict type version of the List struct for better performance.
//Also it has some features to calculate Series the fast possible way, while still saving all the data
but keeping the storange just as much as is needed:
- The constant updated Value is saved in the root and gets added to the data and list when the candle closes
- NILO(New In; Last Out): updates des value to the list and deletes the last one which isnt anymore needed for the calculation
- All data will have the same lenght, so the whole chart information can be saved chronical in a signal array of Pointers to FloatArrays
*/

type Fist struct {
	root Element
	data *[]float64
	len  int //List Lenght, excluding this element
}

func New() *Fist { return new(Fist).init() }

func (l *Fist) init() *Fist {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func (l *Fist) Len() int         { return l.len }
func (l *Fist) Root() float64    { return l.root.value }
func (l *Fist) Update(v float64) { l.root.value = v }
func (l *Fist) Last() float64    { return l.root.prev.value }

//Get gets you the value from the data array. It does not check for errors!
func (l *Fist) GetData(index int) float64 {
	return (*l.data)[index]
}

//Gets you the value from the stack.
func (l *Fist) GetList(index int) float64 {
	if index >= l.len {
		return 0
	}
	e := l.root.next
	for i := 0; i < index; i++ {
		e = e.next
	}
	return e.value
}

func (l *Fist) Push(v float64) {
	o := l.root.next
	e := Element{value: v, prev: &l.root, next: o}
	o.prev = &e
	l.root.next = &e
	l.len++
}

func (l *Fist) NILO() float64 {
	l.Push(l.Root())
	return l.PopLast()
}

func (l *Fist) PopLast() (v float64) {
	e := l.root.prev
	v = e.value
	l.root.prev = e.prev
	e.prev.next = e.next
	e.next = nil
	e.prev = nil
	l.len--
	return
}
