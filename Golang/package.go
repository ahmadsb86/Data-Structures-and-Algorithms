package main

import (
	"fmt"
)


// =========================================================
// ===================== Linked List =======================
// =========================================================

type LLNode[T any] struct {
	val T
	nxt *LLNode[T]	//next LLNode
	prv *LLNode[T]	//previous LLNode
}

type LinkedList[T any] struct{
	head *LLNode[T]	//For remembering where to start traversing form
	tail *LLNode[T]	//For appending LLNodes to the tail
	idx *LLNode[T]	//Index LLNode for traversing
	size int
	Traversed bool
}


func (ll *LinkedList[T]) Push(input T) {
	alc := &LLNode[T]{val: input, prv:ll.tail}	//allocate mem for new node
	if ll.head == nil {
		ll.head = alc
		ll.tail = ll.head
	} else {
		ll.tail.nxt = alc
		ll.tail = alc
	}
	ll.size++
}

func (ll *LinkedList[T]) Pop() T{
	if(ll.size == 0){
		panic("Cannot pop from empty linked list")
	}
	elm := ll.tail.val
	if ll.size == 1 {
		ll.head = nil
		ll.tail = nil
	} else{
		toPop := ll.tail
		ll.tail = ll.tail.prv
		toPop.prv.nxt = nil
	}
	ll.size --
	return elm
}

func (ll *LinkedList[T]) Shift(input T) {
	alc := &LLNode[T]{val: input, nxt:ll.head}	//allocate mem for new node
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
	elm := ll.head.val
	if ll.size == 1 {
		ll.head = nil
		ll.tail = nil
	} else{
		toPop := ll.head
		ll.head = ll.head.nxt
		toPop.nxt.prv = nil
	}
	ll.size --
	return elm
}

func (ll *LinkedList[T]) Display(params ...string){	//utility func for displaying LL
	//TODO make optional paramater thingy better
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
	return ll.idx.val
}

func (ll *LinkedList[T]) Traverse() T{
	var x T
	if(ll.idx.nxt!=nil){
		ll.idx = ll.idx.nxt
		x = ll.idx.val
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
	return ll.idx.val
}

func (ll *LinkedList[T]) ReverseTraverse() T{
	var x T
	if(ll.idx.prv!=nil){
		ll.idx = ll.idx.prv
		x = ll.idx.val
	}else{
		ll.Traversed = true
	}
	return x 
}


// =========================================================
// ======================== Stack ==========================
// =========================================================


type Stack[T any] struct {	//Built using Linked List class
	lst LinkedList[T]
}

func (stack *Stack[T]) Peek() T {
	if stack.lst.size == 0 {
		panic("Cannot Peek empty stack")
	}
	return stack.lst.tail.val
}

func (stack *Stack[T]) Push(elm T) {
	stack.lst.Push(elm)
}

func (stack *Stack[T]) Pop() (elm T) {
	return stack.lst.Pop()
}

// =========================================================
// ======================== Queue ==========================
// =========================================================


type Queue[T any] struct {	//Built using Linked List class
	lst LinkedList[T]
}

func (queue *Queue[T]) Peek() T {
	if queue.lst.size == 0 {
		panic("Cannot Peek empty stack")
	}
	return queue.lst.head.val
}

func (queue *Queue[T]) Push(elm T) {
	queue.lst.Push(elm)
}

func (queue *Queue[T]) Pop() (elm T) {
	return queue.lst.Slice()
}


// =========================================================
// ======================= Hashmap =========================
// =========================================================

//Built-in Implementation with 'map'



// =========================================================
// =================== Priority Queue ======================
// =========================================================

//TODO finish implementation

type PQNode[T any] struct{
	val int
	data T
	parent *PQNode[T]
	left *PQNode[T]
	right *PQNode[T]
}

type PriorityQueue[T any] struct{
	root *PQNode[T]
	tail *PQNode[T]
	size int
}

func (pq *PriorityQueue[T]) Peek() T{
	return pq.root.data
}

func (pq *PriorityQueue[T]) Push(elm T, priority int){
	alc := &PQNode[T]{data: elm, parent: pq.tail, val: priority}
	if pq.root == nil{
		pq.root = alc
		pq.tail = alc
	}else{
		if(pq.size%2 == 0){
			pq.tail.left = alc
		}else{
			pq.tail.right = alc
		}
		pq.tail = alc
	}
}

// func (pq *PriorityQueue[]) Display(){
// }

// func InitTraverse(pq *PriorityQueue[]) 


// =========================================================
// ======================== Test ===========================
// =========================================================


func main(){
	pq := PriorityQueue[string]{}
	pq.Push("Hello", 1)
	pq.Push("World", 2)
}