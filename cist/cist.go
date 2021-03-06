package cist

type Element struct {
	//An Elements save all the data which is needed to calculate the Indicators
	//next value!
	next, prev *Element
	value      []float64
}

type Cist struct {
	/*the root Element saves the data which gets constantly updated
	When a candle closes, these Element gets pushed to the front
	while the Back element(which isnt needed anymore) get deleted
	*/
	root *Element
	data []float64
	len  int
}

//Init a new List
func New() *Cist { return new(Cist).init() }

/*Cist is a strict type version of the List struct.
I created this one to be faster. Type checking would be way to slow and also
the way we need calculations can be boosted a lot.
if you want to create own indicators, check the RSI creation as an example
*/

func (l *Cist) init() *Cist {
	l.root = &Element{}
	l.root.next = l.root
	l.root.prev = l.root
	l.len = 0
	return l
}

func (l *Cist) Len() int { return l.len }

//Get the root(actual candle) value
func (l *Cist) Root() []float64 { return l.root.value }

//Get the first value. In pinescript terms sma[1]
func (l *Cist) First() []float64 { return l.root.next.value }

//Getting the last
func (l *Cist) Last() []float64 { return l.root.prev.value }

/*Get gets you the value from the data array.
It does not check for errors, again for performance reasons. Be really
carefully that you don't do mistakes or your program crashes.
Also we get our data in inverse way than it is saved
So Get(0) returns the top or actual value
*/
func (l *Cist) Get(index int) float64 {
	return l.data[l.len-index-1]
}

//GetData return the whole list data
func (l *Cist) GetData() []float64 {
	return l.data
}

//AddToData float values to the array
func (l *Cist) AddToData(f float64) {
	l.data = append((l.data), f)
	l.len = l.len + 1
}

//InitData fills the Data array
func (l *Cist) InitData(f []float64) {
	l.data = f
	l.len = len(f)
}

//Push adds a new element in the front of the list
func (l *Cist) Push(v ...float64) {
	o := l.root.next
	e := Element{value: v, prev: l.root, next: o}
	o.prev = &e
	l.root.next = &e
}

//Deletes last list element
func (l *Cist) PopLast() *Element {
	e := l.root.prev
	l.root.prev = e.prev
	e.prev.next = e.next
	e.next = nil
	e.prev = nil
	return e
}

//GetEle gets you an Element
func (l *Cist) GetEle(index int) (e *Element) {
	e = l.root.next
	for i := 0; i < index; i++ {
		e = e.next
	}
	return e
}

//Prev Gets you the element before an element
func (e *Element) Prev() *Element {
	return e.prev
}

//Next Gets you the next element
func (e *Element) Next() *Element {
	return e.next
}

//Init the Elements with the data needed for a fast calculation
func (l *Cist) FillElements(f ...[]float64) {
	le := len(f[0])
	for i := 0; i < le; i++ {
		var temp []float64
		for _, v := range f {
			temp = append(temp, v[i])
		}
		l.Push(temp...)
	}
}

//NILO New In Last Out. Pushes the root if no input
func (l *Cist) NILO(f ...float64) *Element {
	if len(f) == 0 {
		l.Push(l.Root()...)
	} else {
		l.Push(f...)
	}
	return l.PopLast()
}

//Updates the Root Element and data
func (l *Cist) Update(indicator float64, element ...float64) {
	l.data[l.len-1] = indicator
	l.root.value = element
}

//Add pushes the
func (l *Cist) Add() {
	l.NILO()
	l.AddToData(l.data[l.len-1])
}

func (l *Cist) Lowest(pos int) *Element {
	lowest := l.root.next.Get(pos)
	lowElement := l.root.next
	i := lowElement.next
	for i != l.root {
		if i.Get(pos) < lowest {
			lowElement = i
			lowest = i.Get(pos)
		}
		i = i.next
	}
	return lowElement
}

func (l *Cist) Highest(pos int) *Element {
	highest := l.root.next.Get(pos)
	highElement := l.root.next
	i := highElement.next

	for i != l.root {
		if i.Get(pos) > highest {
			highElement = i
			highest = i.Get(pos)
		}
		i = i.next
	}
	return highElement
}

func (e *Element) Get(pos int) float64 {
	return e.value[pos]
}

func (e *Element) Value() []float64 {
	return e.value
}
