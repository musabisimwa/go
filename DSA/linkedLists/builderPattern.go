package main

import "fmt"

// Node repr
type Node[T any] struct {
	data T
	next *Node[T]
}

// LinkedList
type LinkedList[T any] struct {
	head *Node[T]
}

// LinkedListBuilder to be used for the builder pattern.
type LinkedListBuilder[T any] struct {
	list *LinkedList[T]
}

//  creates a new LinkedListBuilder.
func NewLinkedListBuilder[T any]() *LinkedListBuilder[T] {
	return &LinkedListBuilder[T]{list: &LinkedList[T]{}}
}


func (builder *LinkedListBuilder[T]) Insert(data T) *LinkedListBuilder[T] {
	newNode := &Node[T]{data: data}
	if builder.list.head == nil {
		builder.list.head = newNode
	} else {
		current := builder.list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	return builder
}


func (builder *LinkedListBuilder[T]) Build() *LinkedList[T] {
	return builder.list
}

func (l *LinkedList[T]) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := NewLinkedListBuilder[string]().
		Insert("yet").
		Insert("another").
		Insert("...").
		Build()

	list.Print()
}
