package main

import "fmt"

// https://go.dev/tour/generics/2

type ListElement[T any] struct {
	next  *ListElement[T]
	value T
}

type SingleLinkedList[T any] struct {
	start            *ListElement[T]
	end              *ListElement[T]
	numberOfElements int
}

func (list SingleLinkedList[any]) String() string {
	returnValue := "["
	currentElement := list.start
	for i := 0; i < list.numberOfElements; i++ {
		if i > 0 {
			returnValue += ", "
		}
		returnValue += fmt.Sprintf("%v", currentElement.value)
		currentElement = currentElement.next
	}
	returnValue += "]"
	return returnValue
}

func Add[T any](list *SingleLinkedList[T], newElement T) {
	newListElement := ListElement[T]{next: nil, value: newElement}
	endListElement := list.end

	if endListElement != nil {
		endListElement.next = &newListElement
	} else {
		list.start = &newListElement
	}

	list.end = &newListElement
	list.numberOfElements++
}

func main() {
	list := SingleLinkedList[int]{start: nil, end: nil, numberOfElements: 0}
	fmt.Println(list)
	Add(&list, 1)
	Add(&list, 2)
	Add(&list, 13)
	fmt.Println(list)
}
