package main

import (
	"fmt"
)

//Defines two important things, the size of the list and the last element
type List struct {
	last *Element
	size int
}

//Defines the type of the element and make a pointer to the next
type Element struct {
	value interface{}
	next  *Element
}

//Return the list length
func (l *List) Length() int {
	return l.size
}

//Inserts a new element on the list
func (l *List) Insert(value interface{}) {
	l.last = &Element{value, l.last}
	l.size++
}

// Remove the last element from the list and return it's value
// If the list is empty, return nil
func (l *List) Delete() (value interface{}) {
	if l.size > 0 {
		value, l.last = l.last.value, l.last.next
		l.size--
		return
	}
	return nil
}

func main() {
	list := new(List)

	list.Insert("Foo")
	list.Insert("Bar")
	list.Insert("Mah nigga")
	list.Insert("Dayuuum son!")

	fmt.Println("List length: ", list.Length(), "\n")

	for list.Length() > 0 {
		fmt.Println(list.Delete().(string))
	}

	fmt.Println("\nList length: ", list.Length())
}