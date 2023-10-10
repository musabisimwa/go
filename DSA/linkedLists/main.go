package main

import "fmt"
// Linked list

type Node[T any] struct{
	data T 
	next *Node[T]
}

type LinkedList[T any] struct{
	head *Node[T]
}

//insert required
func(l *LinkedList[T]) insert(data T){
	d:= &Node[T]{
		data:data,
	}
	if l.head ==nil{
		l.head = d
	}else{
		current:= l.head
		for current.next !=nil{
			current =current.next 
		}
		current.next = d
	}
}

func(l *LinkedList[T])Print(){
	current:= l.head
	for current!=nil{
		fmt.Printf("%v -> ",current.data)
		current = current.next
	}
	fmt.Println("nil")

}


func main() {

	list := LinkedList[string]{}
	list.insert("yet")
	list.insert("another")
	list.insert("...")
	list.Print()
	
}