package main

import (
	"fmt"
)

// =========================================================
// ===================== Linked List =======================
// =========================================================

type node[T any] struct {
	val T
	nxt *node[T]	//next node
	prv *node[T]	//previous node
}
type LinkedList[T any] struct{
	head *node[T]	//For remembering where to start traversing form
	tail *node[T]	//For appending nodes to the tail
	idx *node[T]	//Index node for traversing
	size int
	Traversed bool
}


func (ll *LinkedList[T]) Push(input T) {
	alc := &node[T]{val: input, prv:ll.tail}	//allocate mem for new node
	if ll.head == nil {
		ll.head = alc
		ll.tail = ll.head
	} else {
		(*ll.tail).nxt = alc
		ll.tail = alc
	}
	ll.size++
}

func (ll *LinkedList[T]) Pop() T{
	if(ll.size == 0){
		panic("Cannot pop from empty linked list")
	}
	elm := (*ll.tail).val
	if ll.size == 1 {
		ll.head = nil
		ll.tail = nil
	} else{
		toPop := ll.tail
		ll.tail = ll.tail.prv
		(*(*toPop).prv).nxt = nil
	}
	ll.size --
	return elm
	// toPop, elm := ll.tail, (*ll.tail).val
	// ll.tail = ll.tail.prv
	// (*(*toPop).prv).nxt = nil
	// ll.size --
	// return elm
}

func (ll *LinkedList[T]) Shift(input T) {
	alc := &node[T]{val: input, nxt:ll.head}	//allocate mem for new node
	if ll.head == nil {
		ll.head = alc
		ll.tail = ll.head
	} else {
		(*ll.head).prv = alc
		ll.head = alc
	}
	ll.size++
}

func (ll *LinkedList[T]) Slice() T{
	if(ll.size == 0){
		panic("Cannot Slice from empty linked list")
	}
	elm := (*ll.head).val
	if ll.size == 1 {
		ll.head = nil
		ll.tail = nil
	} else{
		toPop := ll.head
		ll.head = ll.head.nxt
		(*(*toPop).nxt).prv = nil
	}
	ll.size --
	return elm
}

func (ll *LinkedList[T]) Display(params ...string){	//utility func for displaying LL
	if(ll.size == 0){
		return
	}
	divider := " "
	if(len(params)!=0){
		divider = params[0]
	}
	for i:=ll.InitTraverse(); !ll.Traversed; i=ll.Traverse(){
		fmt.Printf("%v%s", i, divider)
	}
}

func (ll *LinkedList[T]) InitTraverse() T{
	if(ll.size == 0){
		panic("Cannot traverse empty linked list")
	}
	ll.Traversed = false
	ll.idx = ll.head
	return (*ll.idx).val
}

func (ll *LinkedList[T]) Traverse() T{
	var x T
	if((*ll.idx).nxt!=nil){
		ll.idx = (*ll.idx).nxt
		x = (*ll.idx).val
	}else{
		ll.Traversed = true
	}
	return x 
}

func (ll *LinkedList[T]) InitReverseTraverse() T{
	if(ll.size == 0){
		panic("Cannot traverse empty linked list")
	}
	ll.Traversed = false
	ll.idx = ll.tail
	return (*ll.idx).val
}

func (ll *LinkedList[T]) ReverseTraverse() T{
	var x T
	if((*ll.idx).prv!=nil){
		ll.idx = (*ll.idx).prv
		x = (*ll.idx).val
	}else{
		ll.Traversed = true
	}
	return x 
}


// =========================================================
// ======================== Stack ==========================
// =========================================================


type Stack[T any] struct {	//Built on top of Linked List class
	lst LinkedList[T]
}

func (stack *Stack[T]) Peek() T {
	if stack.lst.size == 0 {
		panic("Cannot Peek empty stack")
	}
	return (*stack.lst.tail).val
}

func (stack *Stack[T]) Push(elm T) {
	stack.lst.Push(elm)
}

func (stack *Stack[T]) Pop() (elm T) {
	return stack.lst.Pop()
}


// =========================================================
// ======================== Test ===========================
// =========================================================


func main(){
	lst := LinkedList[int]{}
	lst.Push(1)
	lst.Push(2)
	lst.Push(3)
	lst.Pop()
	lst.Pop()
	lst.Pop()
	lst.Display()
}